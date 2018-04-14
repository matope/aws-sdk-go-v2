package s3_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type testErrorCase struct {
	RespFn           func() *http.Response
	ReqID, HostID    string
	Code, Msg        string
	Extra            map[string]string
	WithoutStatusMsg bool
}

var testUnmarshalCases = []testErrorCase{
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode: 301,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"abc123"},
					"X-Amz-Id-2":       []string{"321cba"},
				},
				Body:          ioutil.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			}
		},
		ReqID:  "abc123",
		HostID: "321cba",
		Code:   "BucketRegionError", Msg: "incorrect region, the bucket is not in 'mock-region' region",
		Extra: map[string]string{},
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode: 403,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"abc123"},
					"X-Amz-Id-2":       []string{"321cba"},
				},
				Body:          ioutil.NopCloser(bytes.NewReader(nil)),
				ContentLength: 0,
			}
		},
		ReqID:  "abc123",
		HostID: "321cba",
		Code:   "Forbidden", Msg: "Forbidden",
		Extra: map[string]string{},
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode: 400,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"abc123"},
					"X-Amz-Id-2":       []string{"321cba"},
				},
				Body:          ioutil.NopCloser(bytes.NewReader(nil)),
				ContentLength: 0,
			}
		},
		ReqID:  "abc123",
		HostID: "321cba",
		Code:   "BadRequest", Msg: "Bad Request",
		Extra: map[string]string{},
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode: 404,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"abc123"},
					"X-Amz-Id-2":       []string{"321cba"},
				},
				Body:          ioutil.NopCloser(bytes.NewReader(nil)),
				ContentLength: 0,
			}
		},
		ReqID:  "abc123",
		HostID: "321cba",
		Code:   "NotFound", Msg: "Not Found",
		Extra: map[string]string{},
	},
	{
		// SDK only reads request ID and host ID from the header. The values
		// in message body are ignored.
		RespFn: func() *http.Response {
			body := `<Error><Code>SomeException</Code><Message>Exception message</Message><RequestId>ignored-request-id</RequestId><HostId>ignored-host-id</HostId></Error>`
			return &http.Response{
				StatusCode: 500,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"taken-request-id"},
					"X-Amz-Id-2":       []string{"taken-host-id"},
				},
				Body:          ioutil.NopCloser(strings.NewReader(body)),
				ContentLength: int64(len(body)),
			}
		},
		ReqID:  "taken-request-id",
		HostID: "taken-host-id",
		Code:   "SomeException", Msg: "Exception message",
		Extra: map[string]string{
			"RequestId": "ignored-request-id",
			"HostId":    "ignored-host-id",
		},
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode: 404,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"abc123"},
					"X-Amz-Id-2":       []string{"321cba"},
				},
				Body:          ioutil.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			}
		},
		ReqID:  "abc123",
		HostID: "321cba",
		Code:   "NotFound", Msg: "Not Found", WithoutStatusMsg: true,
		Extra: map[string]string{},
	},
	{
		RespFn: func() *http.Response {
			return &http.Response{
				StatusCode: 404,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"abc123"},
					"X-Amz-Id-2":       []string{"321cba"},
				},
				Body:          ioutil.NopCloser(bytes.NewReader(nil)),
				ContentLength: -1,
			}
		},
		ReqID:  "abc123",
		HostID: "321cba",
		Code:   "NotFound", Msg: "Not Found",
		Extra: map[string]string{},
	},
	{
		// Test for Extra elements.
		RespFn: func() *http.Response {
			body := `<Error><Code>SomeException</Code><Message>Exception message</Message><RequestId>xml-request-id</RequestId><HostId>xml-host-id</HostId><Extra1Name>Extra1&#xA;Content</Extra1Name><Extra2Name>Extra2 Content</Extra2Name></Error>`
			return &http.Response{
				StatusCode: 500,
				Header: http.Header{
					"X-Amz-Request-Id": []string{"taken-request-id"},
					"X-Amz-Id-2":       []string{"taken-host-id"},
				},
				Body:          ioutil.NopCloser(strings.NewReader(body)),
				ContentLength: int64(len(body)),
			}
		},
		ReqID:  "taken-request-id",
		HostID: "taken-host-id",
		Code:   "SomeException",
		Msg:    "Exception message",
		Extra: map[string]string{
			"RequestId":  "xml-request-id",
			"HostId":     "xml-host-id",
			"Extra1Name": "Extra1\nContent",
			"Extra2Name": "Extra2 Content",
		},
	},
}

func TestUnmarshalError(t *testing.T) {
	for i, c := range testUnmarshalCases {
		s := s3.New(unit.Config())
		s.Handlers.Send.Clear()
		s.Handlers.Send.PushBack(func(r *request.Request) {
			r.HTTPResponse = c.RespFn()
			if !c.WithoutStatusMsg {
				r.HTTPResponse.Status = fmt.Sprintf("%d%s",
					r.HTTPResponse.StatusCode,
					http.StatusText(r.HTTPResponse.StatusCode))
			}
		})
		req := s.PutBucketAclRequest(&s3.PutBucketAclInput{
			Bucket: aws.String("bucket"), ACL: s3.BucketCannedACLPublicRead,
		})
		_, err := req.Send()

		if err == nil {
			t.Fatalf("%d, expected error, got nil", i)
		}
		if e, a := c.Code, err.(awserr.Error).Code(); e != a {
			t.Errorf("%d, Code: expect %s, got %s", i, e, a)
		}
		if e, a := c.Msg, err.(awserr.Error).Message(); e != a {
			t.Errorf("%d, Message: expect %s, got %s", i, e, a)
		}
		if e, a := c.ReqID, err.(awserr.RequestFailure).RequestID(); e != a {
			t.Errorf("%d, RequestID: expect %s, got %s", i, e, a)
		}
		if e, a := c.HostID, err.(s3.RequestFailure).HostID(); e != a {
			t.Errorf("%d, HostID: expect %s, got %s", i, e, a)
		}
		if e, a := c.Extra, err.(s3.RequestFailure).Extra(); !reflect.DeepEqual(e, a) {
			t.Errorf("%d, Extra: expect %+v, got %+v", i, e, a)
		}
	}
}

const completeMultiResp = `
163
<?xml version="1.0" encoding="UTF-8"?>

<CompleteMultipartUploadResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/"><Location>https://bucket.s3-us-west-2.amazonaws.com/key</Location><Bucket>bucket</Bucket><Key>key</Key><ETag>&quot;a7d414b9133d6483d9a1c4e04e856e3b-2&quot;</ETag></CompleteMultipartUploadResult>
0
`

func Test200NoErrorUnmarshalError(t *testing.T) {
	s := s3.New(unit.Config())
	s.Handlers.Send.Clear()
	s.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header: http.Header{
				"X-Amz-Request-Id": []string{"abc123"},
				"X-Amz-Id-2":       []string{"321cba"},
			},
			Body:          ioutil.NopCloser(strings.NewReader(completeMultiResp)),
			ContentLength: -1,
		}
		r.HTTPResponse.Status = http.StatusText(r.HTTPResponse.StatusCode)
	})
	req := s.CompleteMultipartUploadRequest(&s3.CompleteMultipartUploadInput{
		Bucket: aws.String("bucket"), Key: aws.String("key"),
		UploadId: aws.String("id"),
		MultipartUpload: &s3.CompletedMultipartUpload{Parts: []s3.CompletedPart{
			{ETag: aws.String("etag"), PartNumber: aws.Int64(1)},
		}},
	})
	_, err := req.Send()

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
}

const completeMultiErrResp = `<Error><Code>SomeException</Code><Message>Exception message</Message></Error>`

func Test200WithErrorUnmarshalError(t *testing.T) {
	s := s3.New(unit.Config())
	s.Handlers.Send.Clear()
	s.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = &http.Response{
			StatusCode: 200,
			Header: http.Header{
				"X-Amz-Request-Id": []string{"abc123"},
				"X-Amz-Id-2":       []string{"321cba"},
			},
			Body:          ioutil.NopCloser(strings.NewReader(completeMultiErrResp)),
			ContentLength: -1,
		}
		r.HTTPResponse.Status = http.StatusText(r.HTTPResponse.StatusCode)
	})
	req := s.CompleteMultipartUploadRequest(&s3.CompleteMultipartUploadInput{
		Bucket: aws.String("bucket"), Key: aws.String("key"),
		UploadId: aws.String("id"),
		MultipartUpload: &s3.CompletedMultipartUpload{Parts: []s3.CompletedPart{
			{ETag: aws.String("etag"), PartNumber: aws.Int64(1)},
		}},
	})
	_, err := req.Send()

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if e, a := "SomeException", err.(awserr.Error).Code(); e != a {
		t.Errorf("Code: expect %s, got %s", e, a)
	}
	if e, a := "Exception message", err.(awserr.Error).Message(); e != a {
		t.Errorf("Message: expect %s, got %s", e, a)
	}
	if e, a := "abc123", err.(s3.RequestFailure).RequestID(); e != a {
		t.Errorf("RequestID: expect %s, got %s", e, a)
	}
	if e, a := "321cba", err.(s3.RequestFailure).HostID(); e != a {
		t.Errorf("HostID: expect %s, got %s", e, a)
	}
}

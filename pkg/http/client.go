package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/ernesto-jimenez/httplogger"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ClientTimeoutSeconds = 20
)

type HttpClient struct {
	log     *log.Entry
	BaseUrl string
	Auth    Auth
	client  *http.Client
}

type Auth func(req *http.Request)

func BasicAuth(id, secret string) Auth {
	return func(req *http.Request) {
		req.SetBasicAuth(id, secret)
	}
}

func BearerAuth(token string) Auth {
	return func(req *http.Request) {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}

func BearerAuthCustom(prefix, token string) Auth {
	return func(req *http.Request) {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", prefix, token))
	}
}

type httpLogger struct {
	log *log.Entry
}

func newLogger() *httpLogger {
	return &httpLogger{
		log: log.New().WithField("component", "http-client"),
	}
}

func (l *httpLogger) LogRequest(req *http.Request) {
	l.log.Debugf(
		"Request %s %s",
		req.Method,
		req.URL.String(),
	)
}

func (l *httpLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	if err != nil {
		l.log.Error(err)
	} else {
		l.log.Debugf(
			"Response method=%s status=%d durationMs=%d %s",
			req.Method,
			res.StatusCode,
			duration,
			req.URL.String(),
		)
	}
}

func (s *HttpClient) Init() {
	logger := newLogger()
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	s.client = &http.Client{Transport: httplogger.NewLoggedTransport(tr, logger), Timeout: time.Second * ClientTimeoutSeconds}
	s.log = logger.log
}

func (s *HttpClient) FullUrl(path string) string {
	return fmt.Sprintf("%s%s", s.BaseUrl, path)
}

func (s *HttpClient) Request(method string, path string, body []byte, headers map[string]string) ([]byte, error) {
	url := s.FullUrl(path)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	s.log.Infof("Making request: [%s] %s", method, url)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	if s.Auth != nil {
		s.Auth(request)
	}

	return s.performRequest(request)
}

func (s *HttpClient) performRequest(req *http.Request) ([]byte, error) {
	var emptyResponse []byte
	resp, err := s.client.Do(req)
	if err != nil {
		return emptyResponse, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {

		bodyRaw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			s.log.Errorf("Could not decode body")
		} else {
			bodyString := string(bodyRaw)
			s.log.Errorf("Response body = %s", bodyString)
		}

		return bodyRaw, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyResponse, err
	}

	return data, nil
}

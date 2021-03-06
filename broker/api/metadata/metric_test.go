package metadata

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/lindb/lindb/mock"
	"github.com/lindb/lindb/parallel"
)

func TestMetricAPI_getCommonParams(t *testing.T) {
	req, _ := http.NewRequest("GET", "/metadata/suggest", nil)
	_, _, _, err := getCommonParams(req)
	assert.Error(t, err)

	req, _ = http.NewRequest("GET", "/metadata/suggest?db=test", nil)
	db, prefix, limit, err := getCommonParams(req)
	assert.NoError(t, err)
	assert.Equal(t, "test", db)
	assert.Equal(t, "", prefix)
	assert.Equal(t, 100, limit)

	req, _ = http.NewRequest("GET", "/metadata/suggest?db=test&limit=avc", nil)
	_, _, _, err = getCommonParams(req)
	assert.Error(t, err)
}

func TestMetricAPI_SuggestMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := parallel.NewMockExecutorFactory(ctrl)
	exec := parallel.NewMockMetadataExecutor(ctrl)

	factory.EXPECT().NewMetadataBrokerExecutor(gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(exec).AnyTimes()

	api := NewMetricAPI(nil, nil, factory, nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest",
		HandlerFunc:    api.SuggestMetrics,
		ExpectHTTPCode: 500,
	})
	exec.EXPECT().Execute().Return(nil, fmt.Errorf("err"))
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test",
		HandlerFunc:    api.SuggestMetrics,
		ExpectHTTPCode: 500,
	})

	exec.EXPECT().Execute().Return([]string{"a", "b"}, nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test",
		HandlerFunc:    api.SuggestMetrics,
		ExpectHTTPCode: 200,
		ExpectResponse: []string{"a", "b"},
	})
}

func TestMetricAPI_SuggestTagKeys(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := parallel.NewMockExecutorFactory(ctrl)
	exec := parallel.NewMockMetadataExecutor(ctrl)

	factory.EXPECT().NewMetadataBrokerExecutor(gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(exec).AnyTimes()

	api := NewMetricAPI(nil, nil, factory, nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest",
		HandlerFunc:    api.SuggestTagKeys,
		ExpectHTTPCode: 500,
	})
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test",
		HandlerFunc:    api.SuggestTagKeys,
		ExpectHTTPCode: 500,
	})

	exec.EXPECT().Execute().Return(nil, fmt.Errorf("err"))
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test&metric=cpu",
		HandlerFunc:    api.SuggestTagKeys,
		ExpectHTTPCode: 500,
	})

	exec.EXPECT().Execute().Return([]string{"a", "b"}, nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test&metric=cpu",
		HandlerFunc:    api.SuggestTagKeys,
		ExpectHTTPCode: 200,
		ExpectResponse: []string{"a", "b"},
	})
}

func TestMetricAPI_SuggestTagValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := parallel.NewMockExecutorFactory(ctrl)
	exec := parallel.NewMockMetadataExecutor(ctrl)

	factory.EXPECT().NewMetadataBrokerExecutor(gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(exec).AnyTimes()

	api := NewMetricAPI(nil, nil, factory, nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest",
		HandlerFunc:    api.SuggestTagValues,
		ExpectHTTPCode: 500,
	})
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test",
		HandlerFunc:    api.SuggestTagValues,
		ExpectHTTPCode: 500,
	})
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test&metric=cpu",
		HandlerFunc:    api.SuggestTagValues,
		ExpectHTTPCode: 500,
	})

	exec.EXPECT().Execute().Return(nil, fmt.Errorf("err"))
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test&metric=cpu&tagKey=host",
		HandlerFunc:    api.SuggestTagValues,
		ExpectHTTPCode: 500,
	})

	exec.EXPECT().Execute().Return([]string{"a", "b"}, nil)
	mock.DoRequest(t, &mock.HTTPHandler{
		Method:         http.MethodGet,
		URL:            "/metadata/suggest?db=test&metric=cpu&tagKey=host",
		HandlerFunc:    api.SuggestTagValues,
		ExpectHTTPCode: 200,
		ExpectResponse: []string{"a", "b"},
	})
}

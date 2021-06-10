package fetcher

import (
	"github.com/golang/mock/gomock"
	fetcherMock "hyq.com/study/mock"
	"net/http"
	"testing"
)

func TestHttpFetcher_Get_404(t *testing.T) {
	query := "v4/resolve?dn=www.baidu.com&account_id=100195&ip=1.1.1.1&sign=9cc9723684867d" +
		"5b5eb943431220790a&t=1866666666&cuid=xxxxxyyyyyzzzzz&type=dual_stack"

	ctl := gomock.NewController(t)

	defer ctl.Finish()

	mock_resp := new(http.Response)
	mock_resp.StatusCode = 404
	mockFetcher := fetcherMock.NewMockFetcher(ctl)
	mockFetcher.EXPECT().Get(query).Return(mock_resp, nil)
	ServiceCheck(query, mockFetcher)

}

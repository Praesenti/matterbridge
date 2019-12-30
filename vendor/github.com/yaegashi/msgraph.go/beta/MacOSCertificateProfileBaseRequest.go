// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOSCertificateProfileBaseRequestBuilder is request builder for MacOSCertificateProfileBase
type MacOSCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSCertificateProfileBaseRequest
func (b *MacOSCertificateProfileBaseRequestBuilder) Request() *MacOSCertificateProfileBaseRequest {
	return &MacOSCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSCertificateProfileBaseRequest is request for MacOSCertificateProfileBase
type MacOSCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for MacOSCertificateProfileBase
func (r *MacOSCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *MacOSCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSCertificateProfileBase
func (r *MacOSCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *MacOSCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSCertificateProfileBase
func (r *MacOSCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
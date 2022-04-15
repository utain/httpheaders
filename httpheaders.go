// Http header constant
//
// Link https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
package httpheaders

const (
	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.3.2
	Accept string = "Accept"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.3.3
	AcceptCharset string = "Accept-Charset"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.3.4
	AcceptEncoding string = "Accept-Encoding"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.3.5
	AcceptLanguage string = "Accept-Language"

	// RFC 5789
	// https://tools.ietf.org/html/rfc5789#section-3.1
	AcceptPatch string = "Accept-Patch"

	// RFC 7233
	// https://tools.ietf.org/html/rfc7233#section-2.3
	AcceptRanges string = "Accept-Ranges"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlAllowCredentials string = "Access-Control-Allow-Credentials"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlAllowHeaders string = "Access-Control-Allow-Headers"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlAllowMethods string = "Access-Control-Allow-Methods"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlAllowOrigin string = "Access-Control-Allow-Origin"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlExposeHeaders string = "Access-Control-Expose-Headers"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlMaxAge string = "Access-Control-Max-Age"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlRequestHeaders string = "Access-Control-Request-Headers"

	// CORS W3C recommendation for response header field name.
	// https://www.w3.org/TR/cors/
	AccessControlRequestMethod string = "Access-Control-Request-Method"

	// RFC 7234
	// https://tools.ietf.org/html/rfc7234#section-5.1
	Age string = "Age"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-7.4.1
	Allow string = "Allow"

	// RFC 7235
	// https://tools.ietf.org/html/rfc7235#section-4.2
	Authorization string = "Authorization"

	// RFC 7234
	// https://tools.ietf.org/html/rfc7234#section-5.2
	CacheControl string = "Cache-Control"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-6.1
	Connection string = "Connection"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-3.1.2.2
	ContentEncoding string = "Content-Encoding"

	// RFC 6266
	// https://tools.ietf.org/html/rfc6266
	ContentDisposition string = "Content-Disposition"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-3.1.3.2
	ContentLanguage string = "Content-Language"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-3.3.2
	ContentLength string = "Content-Length"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-3.1.4.2
	ContentLocation string = "Content-Location"

	// RFC 7233
	// https://tools.ietf.org/html/rfc7233#section-4.2
	ContentRange string = "Content-Range"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-3.1.1.5
	ContentType string = "Content-Type"

	// RFC 2109
	// https://tools.ietf.org/html/rfc2109#section-4.3.4
	Cookie string = "Cookie"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-7.1.1.2
	Date string = "Date"

	// RFC 7232
	// https://tools.ietf.org/html/rfc7232#section-2.3
	Etag string = "ETag"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.1.1
	Expect string = "Expect"

	// RFC 7234
	// https://tools.ietf.org/html/rfc7234#section-5.3
	Expires string = "Expires"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.5.1
	From string = "From"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-5.4
	Host string = "Host"

	// RFC 7232
	// https://tools.ietf.org/html/rfc7232#section-3.1
	IfMatch string = "If-Match"

	// RFC 7232
	// https://tools.ietf.org/html/rfc7232#section-3.3
	IfModifiedSince string = "If-Modified-Since"

	// RFC 7232
	// https://tools.ietf.org/html/rfc7232#section-3.2
	IfNoneMatch string = "If-None-Match"

	// RFC 7233
	// https://tools.ietf.org/html/rfc7233#section-3.2
	IfRange string = "If-Range"

	// RFC 7232
	// https://tools.ietf.org/html/rfc7232#section-3.4
	IfUnmodifiedSince string = "If-Unmodified-Since"

	// RFC 7232
	// https://tools.ietf.org/html/rfc7232#section-2.2
	LastModified string = "Last-Modified"

	// RFC 5988
	// https://tools.ietf.org/html/rfc5988
	Link string = "Link"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-7.1.2
	Location string = "Location"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.1.2
	MaxForwards string = "Max-Forwards"

	// RFC 6454
	// https://tools.ietf.org/html/rfc6454
	Origin string = "Origin"

	// RFC 7234
	// https://tools.ietf.org/html/rfc7234#section-5.4
	Pragma string = "Pragma"

	// RFC 7235
	// https://tools.ietf.org/html/rfc7235#section-4.3
	ProxyAuthenticate string = "Proxy-Authenticate"

	// RFC 7235
	// https://tools.ietf.org/html/rfc7235#section-4.4
	ProxyAuthorization string = "Proxy-Authorization"

	// RFC 7233
	// https://tools.ietf.org/html/rfc7233#section-3.1
	Range string = "Range"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.5.2
	Referer string = "Referer"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-7.1.3
	RetryAfter string = "Retry-After"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-7.4.2
	Server string = "Server"

	// RFC 2109
	// https://tools.ietf.org/html/rfc2109#section-4.2.2
	SetCookie string = "Set-Cookie"

	// RFC 2965
	// https://tools.ietf.org/html/rfc2965
	SetCookie2 string = "Set-Cookie2"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-4.3
	Te string = "TE"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-4.4
	Trailer string = "Trailer"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-3.3.1
	TransferEncoding string = "Transfer-Encoding"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-6.7
	Upgrade string = "Upgrade"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-5.5.3
	UserAgent string = "User-Agent"

	// RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-7.1.4
	Vary string = "Vary"

	// RFC 7230
	// https://tools.ietf.org/html/rfc7230#section-5.7.1
	Via string = "Via"

	// RFC 7234
	// https://tools.ietf.org/html/rfc7234#section-5.5
	Warning string = "Warning"

	// RFC 7235
	// https://tools.ietf.org/html/rfc7235#section-4.1
	WwwAuthenticate string = "WWW-Authenticate"
)

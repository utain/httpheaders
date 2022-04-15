// Http media type constant
//
// HTTP 1.1: Content-Type and Accept https://tools.ietf.org/html/rfc7231#section-3.1.1.1
package mediatypes

const (
	All                        string = "*/*"
	ApplicationAtomXml         string = "application/atom+xml"
	ApplicationCbor            string = "application/cbor"
	ApplicationFormUrlencoded  string = "application/x-www-form-urlencoded"
	ApplicationJson            string = "application/json"
	ApplicationJsonUtf8        string = "application/json;charset=UTF-8"
	ApplicationOctetStream     string = "application/octet-stream"
	ApplicationPdf             string = "application/pdf"
	ApplicationProblemJson     string = "application/problem+json"
	ApplicationProblemJsonUtf8 string = "application/problem+json;charset=UTF-8"
	ApplicationProblemXml      string = "application/problem+xml"
	ApplicationRssXml          string = "application/rss+xml"
	ApplicationNdjson          string = "application/x-ndjson"
	ApplicationStreamJson      string = "application/stream+json"
	ApplicationXhtmlXml        string = "application/xhtml+xml"
	ApplicationXml             string = "application/xml"
	ImageGif                   string = "image/gif"
	ImageJpeg                  string = "image/jpeg"
	ImagePng                   string = "image/png"
	MultipartFormData          string = "multipart/form-data"
	MultipartMixed             string = "multipart/mixed"
	MultipartRelated           string = "multipart/related"
	TextEventStream            string = "text/event-stream"
	TextHtml                   string = "text/html"
	TextMarkdown               string = "text/markdown"
	TextPlain                  string = "text/plain"
	TextXml                    string = "text/xml"
)

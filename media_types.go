package mediatypes

// MediaType represents a media type.
type MediaType struct {
	// Name is the media type such as "text/plain" or "application/json".
	name string

	// Extensions is a list of file extensions that are associated with this media type.
	extensions []string

	// Format is the format of the media type such as "application/json" or
	// "application/xml", or empty if unknown.
	format string

	// Registered returns true if the media type is registered with IANA.
	registered bool
}

// String returns the media type as a string.
func (m *MediaType) String() string {
	return m.name
}

// Name returns the media type name.
func (m *MediaType) Name() string {
	return m.name
}

// Format returns the media type format.
func (m *MediaType) Format() string {
	return m.format
}

// Extensions returns a list of file extensions that are associated with this media type.
func (m *MediaType) Extensions() []string {
	return m.extensions
}

// Registered returns true if the media type is registered with IANA.
func (m *MediaType) Registered() bool {
	return m.registered
}

// ByExtension returns the media type with the given file extension.
func ByExtension(ext string) []MediaType {
	var result []MediaType
	for _, m := range mediaTypes {
		for _, e := range m.extensions {
			if e == ext {
				result = append(result, m)
			}
		}
	}
	return result
}

// mediaTypes returns a list of all media types.
var mediaTypes = []MediaType{
	{
		name:       "application/1d-interleaved-parityfec",
		registered: true,
	},
	{
		name:       "application/3gpdash-qoe-report+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/3gpp-ims+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/3gppHal+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/3gppHalForms+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/A2L",
		registered: true,
	},
	{
		name:       "application/AML",
		registered: true,
	},
	{
		name:       "application/ATF",
		registered: true,
	},
	{
		name:       "application/ATFX",
		registered: true,
	},
	{
		name:       "application/ATXML",
		registered: true,
	},
	{
		name:       "application/CALS-1840",
		registered: true,
	},
	{
		name:       "application/CDFX+XML",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/CEA",
		registered: true,
	},
	{
		name:       "application/CSTAdata+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/DCD",
		registered: true,
	},
	{
		name:       "application/DII",
		registered: true,
	},
	{
		name:       "application/DIT",
		registered: true,
	},
	{
		name:       "application/EDI-X12",
		registered: true,
	},
	{
		name:       "application/EDI-consent",
		registered: true,
	},
	{
		name:       "application/EDIFACT",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.Comment+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.Control+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.DeviceInfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.LegacyESN+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.ProviderInfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.ServiceInfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.SubscriberInfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.VEDS+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.cap+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/EmergencyCallData.eCall.MSD",
		registered: true,
	},
	{
		name:       "application/IOTP",
		registered: true,
	},
	{
		name:       "application/LXF",
		registered: true,
	},
	{
		name:       "application/MF4",
		registered: true,
	},
	{
		name:       "application/ODA",
		registered: true,
	},
	{
		name:       "application/ODX",
		registered: true,
	},
	{
		name:       "application/PDX",
		registered: true,
	},
	{
		name:       "application/TETRA_ISI",
		registered: true,
	},
	{
		name:       "application/acad",
		registered: false,
		extensions: []string{
			"dwg",
		},
	},
	{
		name:       "application/ace+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/ace+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/activemessage",
		registered: true,
	},
	{
		name:       "application/activity+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/aif+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/aif+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-cdni+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-cdnifilter+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-costmap+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-costmapfilter+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-directory+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-endpointcost+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-endpointcostparams+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-endpointprop+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-endpointpropparams+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-error+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-networkmap+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-networkmapfilter+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-propmap+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-propmapparams+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-updatestreamcontrol+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/alto-updatestreamparams+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/andrew-inset",
		registered: true,
		extensions: []string{
			"ez",
		},
	},
	{
		name:       "application/applefile",
		registered: true,
	},
	{
		name:       "application/applixware",
		registered: false,
		extensions: []string{
			"aw",
		},
	},
	{
		name:       "application/arj",
		registered: false,
		extensions: []string{
			"arj",
		},
	},
	{
		name:       "application/at+jwt",
		format:     "application/jwt",
		registered: true,
	},
	{
		name:       "application/atom+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"atom", "xml",
		},
	},
	{
		name:       "application/atomcat+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"atomcat",
		},
	},
	{
		name:       "application/atomdeleted+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/atomicmail",
		registered: true,
	},
	{
		name:       "application/atomsvc+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"atomsvc",
		},
	},
	{
		name:       "application/atsc-dwd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/atsc-dynamic-event-message",
		registered: true,
	},
	{
		name:       "application/atsc-held+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/atsc-rdt+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/atsc-rsat+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/auth-policy+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/automationml-aml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/automationml-amlx+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/bacnet-xdd+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/base64",
		registered: false,
		extensions: []string{
			"mm", "mme",
		},
	},
	{
		name:       "application/batch-SMTP",
		registered: true,
	},
	{
		name:       "application/beep+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/binhex",
		registered: false,
		extensions: []string{
			"hqx",
		},
	},
	{
		name:       "application/binhex4",
		registered: false,
		extensions: []string{
			"hqx",
		},
	},
	{
		name:       "application/book",
		registered: false,
		extensions: []string{
			"boo", "book",
		},
	},
	{
		name:       "application/calendar+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/calendar+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/call-completion",
		registered: true,
	},
	{
		name:       "application/captive+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/cbor",
		registered: true,
	},
	{
		name:       "application/cbor-seq",
		registered: true,
	},
	{
		name:       "application/cccex",
		registered: true,
	},
	{
		name:       "application/ccmp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/ccxml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"ccxml",
		},
	},
	{
		name:       "application/cda+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/cdf",
		registered: false,
		extensions: []string{
			"cdf",
		},
	},
	{
		name:       "application/cdmi-capability",
		registered: true,
		extensions: []string{
			"cdmia",
		},
	},
	{
		name:       "application/cdmi-container",
		registered: true,
		extensions: []string{
			"cdmic",
		},
	},
	{
		name:       "application/cdmi-domain",
		registered: true,
		extensions: []string{
			"cdmid",
		},
	},
	{
		name:       "application/cdmi-object",
		registered: true,
		extensions: []string{
			"cdmio",
		},
	},
	{
		name:       "application/cdmi-queue",
		registered: true,
		extensions: []string{
			"cdmiq",
		},
	},
	{
		name:       "application/cdni",
		registered: true,
	},
	{
		name:       "application/cea-2018+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/cellml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/cfw",
		registered: true,
	},
	{
		name:       "application/city+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/clariscad",
		registered: false,
		extensions: []string{
			"ccad",
		},
	},
	{
		name:       "application/clr",
		registered: true,
	},
	{
		name:       "application/clue+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/clue_info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/cms",
		registered: true,
	},
	{
		name:       "application/cnrp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/coap-group+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/coap-payload",
		registered: true,
	},
	{
		name:       "application/commonground",
		registered: true,
		extensions: []string{
			"dp",
		},
	},
	{
		name:       "application/concise-problem-details+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/conference-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/cose",
		registered: true,
	},
	{
		name:       "application/cose-key",
		registered: true,
	},
	{
		name:       "application/cose-key-set",
		registered: true,
	},
	{
		name:       "application/cose-x509",
		registered: true,
	},
	{
		name:       "application/cpl+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/csrattrs",
		registered: true,
	},
	{
		name:       "application/csta+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/csvm+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/cu-seeme",
		registered: false,
		extensions: []string{
			"cu", "csm",
		},
	},
	{
		name:       "application/cwl",
		registered: true,
	},
	{
		name:       "application/cwl+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/cwt",
		registered: true,
	},
	{
		name:       "application/cybercash",
		registered: true,
	},
	{
		name:       "application/dash+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/dash-patch+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/dashdelta",
		registered: true,
	},
	{
		name:       "application/davmount+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"davmount",
		},
	},
	{
		name:       "application/dca-rft",
		registered: true,
	},
	{
		name:       "application/dec-dx",
		registered: true,
	},
	{
		name:       "application/dialog-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/dicom",
		registered: true,
	},
	{
		name:       "application/dicom+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/dicom+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/dns",
		registered: true,
	},
	{
		name:       "application/dns+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/dns-message",
		registered: true,
	},
	{
		name:       "application/docbook+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"dbk",
		},
	},
	{
		name:       "application/dots+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/drafting",
		registered: false,
		extensions: []string{
			"drw",
		},
	},
	{
		name:       "application/dskpp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/dsptype",
		registered: false,
		extensions: []string{
			"tsp",
		},
	},
	{
		name:       "application/dssc+der",
		registered: true,
		extensions: []string{
			"dssc",
		},
	},
	{
		name:       "application/dssc+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xdssc",
		},
	},
	{
		name:       "application/dvcs",
		registered: true,
	},
	{
		name:       "application/dxf",
		registered: false,
		extensions: []string{
			"dxf",
		},
	},
	{
		name:       "application/ecmascript",
		registered: true,
		extensions: []string{
			"es", "ecma", "js",
		},
	},
	{
		name:       "application/efi",
		registered: true,
	},
	{
		name:       "application/elm+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/elm+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/emma+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"emma",
		},
	},
	{
		name:       "application/emotionml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/encaprtp",
		registered: true,
	},
	{
		name:       "application/envoy",
		registered: false,
		extensions: []string{
			"evy",
		},
	},
	{
		name:       "application/epp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/epub+zip",
		format:     "application/zip",
		registered: true,
		extensions: []string{
			"epub",
		},
	},
	{
		name:       "application/eshop",
		registered: true,
	},
	{
		name:       "application/example",
		registered: true,
	},
	{
		name:       "application/excel",
		registered: false,
		extensions: []string{
			"xl", "xla", "xlb", "xlc", "xld", "xlk", "xll", "xlm", "xls", "xlt", "xlv", "xlw",
		},
	},
	{
		name:       "application/exi",
		registered: true,
		extensions: []string{
			"exi",
		},
	},
	{
		name:       "application/expect-ct-report+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/express",
		registered: true,
	},
	{
		name:       "application/fastinfoset",
		registered: true,
	},
	{
		name:       "application/fastsoap",
		registered: true,
	},
	{
		name:       "application/fdf",
		registered: true,
	},
	{
		name:       "application/fdt+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/fhir+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/fhir+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/fits",
		registered: true,
	},
	{
		name:       "application/flexfec",
		registered: true,
	},
	{
		name:       "application/font-sfnt",
		registered: true,
	},
	{
		name:       "application/font-tdpfr",
		registered: true,
		extensions: []string{
			"pfr",
		},
	},
	{
		name:       "application/font-woff",
		registered: true,
		extensions: []string{
			"woff",
		},
	},
	{
		name:       "application/fractals",
		registered: false,
		extensions: []string{
			"fif",
		},
	},
	{
		name:       "application/framework-attributes+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/freeloader",
		registered: false,
		extensions: []string{
			"frl",
		},
	},
	{
		name:       "application/futuresplash",
		registered: false,
		extensions: []string{
			"spl",
		},
	},
	{
		name:       "application/geo+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/geo+json-seq",
		registered: true,
	},
	{
		name:       "application/geopackage+sqlite3",
		format:     "application/vnd.sqlite3",
		registered: true,
	},
	{
		name:       "application/geoxacml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/ghostview",
		registered: false,
	},
	{
		name:       "application/gltf-buffer",
		registered: true,
	},
	{
		name:       "application/gml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"gml",
		},
	},
	{
		name:       "application/gnutar",
		registered: false,
		extensions: []string{
			"tgz",
		},
	},
	{
		name:       "application/gpx+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"gpx",
		},
	},
	{
		name:       "application/groupwise",
		registered: false,
		extensions: []string{
			"vew",
		},
	},
	{
		name:       "application/gxf",
		registered: false,
		extensions: []string{
			"gxf",
		},
	},
	{
		name:       "application/gzip",
		registered: true,
	},
	{
		name:       "application/h224",
		registered: false,
	},
	{
		name:       "application/held+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/hl7v2+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/hlp",
		registered: false,
		extensions: []string{
			"hlp",
		},
	},
	{
		name:       "application/hta",
		registered: false,
		extensions: []string{
			"hta",
		},
	},
	{
		name:       "application/http",
		registered: true,
	},
	{
		name:       "application/hyperstudio",
		registered: true,
		extensions: []string{
			"stk",
		},
	},
	{
		name:       "application/i-deas",
		registered: false,
		extensions: []string{
			"unv",
		},
	},
	{
		name:       "application/ibe-key-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/ibe-pkg-reply+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/ibe-pp-data",
		registered: true,
	},
	{
		name:       "application/iges",
		registered: true,
		extensions: []string{
			"iges", "igs",
		},
	},
	{
		name:       "application/im-iscomposing+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/index",
		registered: true,
	},
	{
		name:       "application/index.cmd",
		registered: true,
	},
	{
		name:       "application/index.obj",
		registered: true,
	},
	{
		name:       "application/index.response",
		registered: true,
	},
	{
		name:       "application/index.vnd",
		registered: true,
	},
	{
		name:       "application/inf",
		registered: false,
		extensions: []string{
			"inf",
		},
	},
	{
		name:       "application/inkml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"ink", "inkml",
		},
	},
	{
		name:       "application/internet-property-stream",
		registered: false,
		extensions: []string{
			"acx",
		},
	},
	{
		name:       "application/ipfix",
		registered: true,
		extensions: []string{
			"ipfix",
		},
	},
	{
		name:       "application/ipp",
		registered: true,
	},
	{
		name:       "application/isup",
		registered: false,
	},
	{
		name:       "application/its+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/java",
		registered: false,
		extensions: []string{
			"class",
		},
	},
	{
		name:       "application/java-archive",
		registered: false,
		extensions: []string{
			"jar",
		},
	},
	{
		name:       "application/java-byte-code",
		registered: false,
		extensions: []string{
			"class",
		},
	},
	{
		name:       "application/java-serialized-object",
		registered: false,
		extensions: []string{
			"ser",
		},
	},
	{
		name:       "application/java-vm",
		registered: false,
		extensions: []string{
			"class",
		},
	},
	{
		name:       "application/javascript",
		registered: true,
		extensions: []string{
			"js",
		},
	},
	{
		name:       "application/jf2feed+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/jose",
		registered: true,
	},
	{
		name:       "application/jose+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/jrd+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/jscalendar+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/json",
		registered: true,
		extensions: []string{
			"json",
		},
	},
	{
		name:       "application/json-patch+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/json-seq",
		registered: true,
	},
	{
		name:       "application/jsonml+json",
		format:     "application/json",
		registered: false,
		extensions: []string{
			"jsonml",
		},
	},
	{
		name:       "application/jwk+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/jwk-set+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/jwt",
		registered: true,
	},
	{
		name:       "application/kpml-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/kpml-response+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/ld+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/lgr+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/lha",
		registered: false,
		extensions: []string{
			"lha",
		},
	},
	{
		name:       "application/link-format",
		registered: true,
	},
	{
		name:       "application/linkset",
		registered: true,
	},
	{
		name:       "application/linkset+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/load-control+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/logout+jwt",
		format:     "application/jwt",
		registered: true,
	},
	{
		name:       "application/lost+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"lostxml",
		},
	},
	{
		name:       "application/lostsync+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/lpf+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/lzx",
		registered: false,
		extensions: []string{
			"lzx",
		},
	},
	{
		name:       "application/mac-binary",
		registered: false,
		extensions: []string{
			"bin",
		},
	},
	{
		name:       "application/mac-binhex",
		registered: false,
		extensions: []string{
			"hqx",
		},
	},
	{
		name:       "application/mac-binhex40",
		registered: true,
		extensions: []string{
			"hqx",
		},
	},
	{
		name:       "application/mac-compactpro",
		registered: false,
		extensions: []string{
			"cpt",
		},
	},
	{
		name:       "application/macbinary",
		registered: false,
		extensions: []string{
			"bin",
		},
	},
	{
		name:       "application/macwriteii",
		registered: true,
	},
	{
		name:       "application/mads+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mads",
		},
	},
	{
		name:       "application/manifest+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/marc",
		registered: true,
		extensions: []string{
			"mrc",
		},
	},
	{
		name:       "application/marcxml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mrcx",
		},
	},
	{
		name:       "application/mathematica",
		registered: true,
		extensions: []string{
			"ma", "nb", "mb",
		},
	},
	{
		name:       "application/mathematica-old",
		registered: false,
	},
	{
		name:       "application/mathml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mathml",
		},
	},
	{
		name:       "application/mathml-content+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mathml-presentation+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbedlet",
		registered: false,
		extensions: []string{
			"mbd",
		},
	},
	{
		name:       "application/mbms-associated-procedure-description+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-deregister+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-envelope+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-msk+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-msk-response+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-protection-description+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-reception-report+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-register+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-register-response+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-schedule+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbms-user-service-description+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mbox",
		registered: true,
		extensions: []string{
			"mbox",
		},
	},
	{
		name:       "application/mcad",
		registered: false,
		extensions: []string{
			"mcd",
		},
	},
	{
		name:       "application/media-policy-dataset+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/media_control+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mediaservercontrol+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mscml",
		},
	},
	{
		name:       "application/merge-patch+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/metalink+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"metalink",
		},
	},
	{
		name:       "application/metalink4+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"meta4",
		},
	},
	{
		name:       "application/mets+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mets",
		},
	},
	{
		name:       "application/mikey",
		registered: true,
	},
	{
		name:       "application/mime",
		registered: false,
		extensions: []string{
			"aps",
		},
	},
	{
		name:       "application/mipc",
		registered: true,
	},
	{
		name:       "application/missing-blocks+cbor-seq",
		registered: true,
	},
	{
		name:       "application/mmt-aei+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mmt-usd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mods+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mods",
		},
	},
	{
		name:       "application/moss-keys",
		registered: true,
	},
	{
		name:       "application/moss-signature",
		registered: true,
	},
	{
		name:       "application/mosskey-data",
		registered: true,
	},
	{
		name:       "application/mosskey-request",
		registered: true,
	},
	{
		name:       "application/mp21",
		registered: true,
		extensions: []string{
			"m21", "mp21",
		},
	},
	{
		name:       "application/mp4",
		registered: true,
		extensions: []string{
			"mp4", "m4p", "mp4s",
		},
	},
	{
		name:       "application/mpeg4-generic",
		registered: true,
	},
	{
		name:       "application/mpeg4-iod",
		registered: true,
	},
	{
		name:       "application/mpeg4-iod-xmt",
		registered: true,
	},
	{
		name:       "application/mrb-consumer+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/mrb-publish+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/msaccess",
		registered: false,
		extensions: []string{
			"mdb",
		},
	},
	{
		name:       "application/msc-ivr+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/msc-mixer+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/msonenote",
		registered: false,
		extensions: []string{
			"one", "onetoc2", "onetmp", "onepkg",
		},
	},
	{
		name:       "application/mspowerpoint",
		registered: false,
		extensions: []string{
			"pot", "pps", "ppt", "ppz",
		},
	},
	{
		name:       "application/msword",
		registered: true,
		extensions: []string{
			"doc", "dot", "w6w", "wiz", "word",
		},
	},
	{
		name:       "application/mswrite",
		registered: false,
		extensions: []string{
			"wri",
		},
	},
	{
		name:       "application/mud+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/multipart-core",
		registered: true,
	},
	{
		name:       "application/mxf",
		registered: true,
		extensions: []string{
			"mxf",
		},
	},
	{
		name:       "application/n-quads",
		registered: true,
	},
	{
		name:       "application/n-triples",
		registered: true,
	},
	{
		name:       "application/nasdata",
		registered: true,
	},
	{
		name:       "application/netmc",
		registered: false,
		extensions: []string{
			"mcp",
		},
	},
	{
		name:       "application/news-checkgroups",
		registered: true,
	},
	{
		name:       "application/news-groupinfo",
		registered: true,
	},
	{
		name:       "application/news-message-id",
		registered: false,
	},
	{
		name:       "application/news-transmission",
		registered: true,
	},
	{
		name:       "application/nlsml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/node",
		registered: true,
	},
	{
		name:       "application/nss",
		registered: true,
	},
	{
		name:       "application/oauth-authz-req+jwt",
		format:     "application/jwt",
		registered: true,
	},
	{
		name:       "application/oblivious-dns-message",
		registered: true,
	},
	{
		name:       "application/ocsp-request",
		registered: true,
	},
	{
		name:       "application/ocsp-response",
		registered: true,
	},
	{
		name:       "application/octet-stream",
		registered: true,
		extensions: []string{
			"bin", "dms", "lrf", "mar", "so", "dist", "distz", "pkg", "bpk", "dump", "elc", "a", "arc", "arj", "com", "exe", "lha", "lhx", "lzh", "lzx", "o", "psd", "saveme", "uu", "zoo", "class", "buffer", "deploy", "hqx", "obj", "lib", "zip", "gz", "dmg", "iso",
		},
	},
	{
		name:       "application/odm+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/oebps-package+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"opf",
		},
	},
	{
		name:       "application/ogg",
		registered: true,
		extensions: []string{
			"ogx", "ogg",
		},
	},
	{
		name:       "application/olescript",
		registered: false,
		extensions: []string{
			"axs",
		},
	},
	{
		name:       "application/omdoc+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"omdoc",
		},
	},
	{
		name:       "application/onenote",
		registered: false,
		extensions: []string{
			"onetoc", "onetoc2", "onetmp", "onepkg",
		},
	},
	{
		name:       "application/opc-nodeset+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/oscore",
		registered: true,
	},
	{
		name:       "application/oxps",
		registered: true,
		extensions: []string{
			"oxps",
		},
	},
	{
		name:       "application/p21",
		registered: true,
	},
	{
		name:       "application/p21+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/p2p-overlay+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/parityfec",
		registered: true,
	},
	{
		name:       "application/passport",
		registered: true,
	},
	{
		name:       "application/patch-ops-error+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xer",
		},
	},
	{
		name:       "application/pdf",
		registered: true,
		extensions: []string{
			"pdf",
		},
	},
	{
		name:       "application/pem-certificate-chain",
		registered: true,
	},
	{
		name:       "application/pgp-encrypted",
		registered: true,
		extensions: []string{
			"pgp",
		},
	},
	{
		name:       "application/pgp-keys",
		registered: true,
		extensions: []string{
			"key",
		},
	},
	{
		name:       "application/pgp-signature",
		registered: true,
		extensions: []string{
			"asc", "pgp", "sig",
		},
	},
	{
		name:       "application/pics-rules",
		registered: false,
		extensions: []string{
			"prf",
		},
	},
	{
		name:       "application/pidf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/pidf-diff+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/pkcs-12",
		registered: false,
		extensions: []string{
			"p12",
		},
	},
	{
		name:       "application/pkcs-crl",
		registered: false,
		extensions: []string{
			"crl",
		},
	},
	{
		name:       "application/pkcs10",
		registered: true,
		extensions: []string{
			"p10",
		},
	},
	{
		name:       "application/pkcs12",
		registered: true,
	},
	{
		name:       "application/pkcs7-mime",
		registered: true,
		extensions: []string{
			"p7m", "p7c",
		},
	},
	{
		name:       "application/pkcs7-signature",
		registered: true,
		extensions: []string{
			"p7s",
		},
	},
	{
		name:       "application/pkcs8",
		registered: true,
		extensions: []string{
			"p8",
		},
	},
	{
		name:       "application/pkcs8-encrypted",
		registered: true,
	},
	{
		name:       "application/pkix-attr-cert",
		registered: true,
		extensions: []string{
			"ac",
		},
	},
	{
		name:       "application/pkix-cert",
		registered: true,
		extensions: []string{
			"cer", "crt",
		},
	},
	{
		name:       "application/pkix-crl",
		registered: true,
		extensions: []string{
			"crl",
		},
	},
	{
		name:       "application/pkix-pkipath",
		registered: true,
		extensions: []string{
			"pkipath",
		},
	},
	{
		name:       "application/pkixcmp",
		registered: true,
		extensions: []string{
			"pki",
		},
	},
	{
		name:       "application/plain",
		registered: false,
		extensions: []string{
			"text",
		},
	},
	{
		name:       "application/pls+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"pls",
		},
	},
	{
		name:       "application/poc-settings+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/postscript",
		registered: true,
		extensions: []string{
			"ai", "eps", "ps",
		},
	},
	{
		name:       "application/powerpoint",
		registered: false,
		extensions: []string{
			"ppt",
		},
	},
	{
		name:       "application/ppsp-tracker+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/pro_eng",
		registered: false,
		extensions: []string{
			"part", "prt",
		},
	},
	{
		name:       "application/problem+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/problem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/provenance+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/prs.alvestrand.titrax-sheet",
		registered: true,
	},
	{
		name:       "application/prs.cww",
		registered: true,
		extensions: []string{
			"cww",
		},
	},
	{
		name:       "application/prs.cyn",
		registered: true,
	},
	{
		name:       "application/prs.hpub+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/prs.nprend",
		registered: true,
	},
	{
		name:       "application/prs.plucker",
		registered: true,
	},
	{
		name:       "application/prs.rdf-xml-crypt",
		registered: true,
	},
	{
		name:       "application/prs.xsf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/pskc+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"pskcxml",
		},
	},
	{
		name:       "application/pvd+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/qsig",
		registered: false,
	},
	{
		name:       "application/raptorfec",
		registered: true,
	},
	{
		name:       "application/rar",
		registered: false,
		extensions: []string{
			"rar",
		},
	},
	{
		name:       "application/rdap+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/rdf+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"rdf",
		},
	},
	{
		name:       "application/reginfo+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"rif",
		},
	},
	{
		name:       "application/relax-ng-compact-syntax",
		registered: true,
		extensions: []string{
			"rnc",
		},
	},
	{
		name:       "application/remote-printing",
		registered: true,
	},
	{
		name:       "application/reputon+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/resource-lists+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"rl",
		},
	},
	{
		name:       "application/resource-lists-diff+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"rld",
		},
	},
	{
		name:       "application/rfc+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/ringing-tones",
		registered: false,
		extensions: []string{
			"rng",
		},
	},
	{
		name:       "application/riscos",
		registered: true,
	},
	{
		name:       "application/rlmi+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/rls-services+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"rs",
		},
	},
	{
		name:       "application/route-apd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/route-s-tsid+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/route-usd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/rpki-checklist",
		registered: true,
	},
	{
		name:       "application/rpki-ghostbusters",
		registered: true,
		extensions: []string{
			"gbr",
		},
	},
	{
		name:       "application/rpki-manifest",
		registered: true,
		extensions: []string{
			"mft",
		},
	},
	{
		name:       "application/rpki-publication",
		registered: true,
	},
	{
		name:       "application/rpki-roa",
		registered: true,
		extensions: []string{
			"roa",
		},
	},
	{
		name:       "application/rpki-updown",
		registered: true,
	},
	{
		name:       "application/rsd+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"rsd",
		},
	},
	{
		name:       "application/rss+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"rss", "xml",
		},
	},
	{
		name:       "application/rtf",
		registered: true,
		extensions: []string{
			"rtf", "rtx",
		},
	},
	{
		name:       "application/rtploopback",
		registered: true,
	},
	{
		name:       "application/rtx",
		registered: true,
	},
	{
		name:       "application/samlassertion+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/samlmetadata+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/sarif+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/sarif-external-properties+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/sbe",
		registered: true,
	},
	{
		name:       "application/sbml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"sbml",
		},
	},
	{
		name:       "application/scaip+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/scim+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/scvp-cv-request",
		registered: true,
		extensions: []string{
			"scq",
		},
	},
	{
		name:       "application/scvp-cv-response",
		registered: true,
		extensions: []string{
			"scs",
		},
	},
	{
		name:       "application/scvp-vp-request",
		registered: true,
		extensions: []string{
			"spq",
		},
	},
	{
		name:       "application/scvp-vp-response",
		registered: true,
		extensions: []string{
			"spp",
		},
	},
	{
		name:       "application/sdp",
		registered: true,
		extensions: []string{
			"sdp",
		},
	},
	{
		name:       "application/sea",
		registered: false,
		extensions: []string{
			"sea",
		},
	},
	{
		name:       "application/secevent+jwt",
		format:     "application/jwt",
		registered: true,
	},
	{
		name:       "application/senml+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/senml+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/senml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/senml-etch+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/senml-etch+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/senml-exi",
		registered: true,
	},
	{
		name:       "application/sensml+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/sensml+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/sensml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/sensml-exi",
		registered: true,
	},
	{
		name:       "application/sep+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/sep-exi",
		registered: true,
	},
	{
		name:       "application/session-info",
		registered: true,
	},
	{
		name:       "application/set",
		registered: false,
		extensions: []string{
			"set",
		},
	},
	{
		name:       "application/set-payment",
		registered: true,
	},
	{
		name:       "application/set-payment-initiation",
		registered: true,
		extensions: []string{
			"setpay",
		},
	},
	{
		name:       "application/set-registration",
		registered: true,
	},
	{
		name:       "application/set-registration-initiation",
		registered: true,
		extensions: []string{
			"setreg",
		},
	},
	{
		name:       "application/sgml",
		registered: false,
	},
	{
		name:       "application/sgml-open-catalog",
		registered: true,
	},
	{
		name:       "application/shf+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"shf",
		},
	},
	{
		name:       "application/sieve",
		registered: true,
	},
	{
		name:       "application/simple-filter+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/simple-message-summary",
		registered: true,
	},
	{
		name:       "application/simpleSymbolContainer",
		registered: true,
	},
	{
		name:       "application/sipc",
		registered: true,
	},
	{
		name:       "application/sla",
		registered: false,
		extensions: []string{
			"stl",
		},
	},
	{
		name:       "application/slate",
		registered: true,
	},
	{
		name:       "application/smil",
		registered: true,
		extensions: []string{
			"smi", "smil",
		},
	},
	{
		name:       "application/smil+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"smi", "smil",
		},
	},
	{
		name:       "application/smpte336m",
		registered: true,
	},
	{
		name:       "application/soap+fastinfoset",
		registered: true,
	},
	{
		name:       "application/soap+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/solids",
		registered: false,
		extensions: []string{
			"sol",
		},
	},
	{
		name:       "application/sounder",
		registered: false,
		extensions: []string{
			"sdr",
		},
	},
	{
		name:       "application/sparql-query",
		registered: true,
		extensions: []string{
			"rq",
		},
	},
	{
		name:       "application/sparql-results+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"srx",
		},
	},
	{
		name:       "application/spdx+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/spirits-event+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/sql",
		registered: true,
	},
	{
		name:       "application/srgs",
		registered: true,
		extensions: []string{
			"gram",
		},
	},
	{
		name:       "application/srgs+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"grxml",
		},
	},
	{
		name:       "application/sru+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"sru",
		},
	},
	{
		name:       "application/ssdl+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"ssdl",
		},
	},
	{
		name:       "application/ssml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"ssml",
		},
	},
	{
		name:       "application/step",
		registered: false,
		extensions: []string{
			"step", "stp",
		},
	},
	{
		name:       "application/stix+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/streamingmedia",
		registered: false,
		extensions: []string{
			"ssm",
		},
	},
	{
		name:       "application/swid+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/swid+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/tamp-apex-update",
		registered: true,
	},
	{
		name:       "application/tamp-apex-update-confirm",
		registered: true,
	},
	{
		name:       "application/tamp-community-update",
		registered: true,
	},
	{
		name:       "application/tamp-community-update-confirm",
		registered: true,
	},
	{
		name:       "application/tamp-error",
		registered: true,
	},
	{
		name:       "application/tamp-sequence-adjust",
		registered: true,
	},
	{
		name:       "application/tamp-sequence-adjust-confirm",
		registered: true,
	},
	{
		name:       "application/tamp-status-query",
		registered: true,
	},
	{
		name:       "application/tamp-status-response",
		registered: true,
	},
	{
		name:       "application/tamp-update",
		registered: true,
	},
	{
		name:       "application/tamp-update-confirm",
		registered: true,
	},
	{
		name:       "application/taxii+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/td+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/tei+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"tei", "teicorpus",
		},
	},
	{
		name:       "application/thraud+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"tfi",
		},
	},
	{
		name:       "application/timestamp-query",
		registered: true,
	},
	{
		name:       "application/timestamp-reply",
		registered: true,
	},
	{
		name:       "application/timestamped-data",
		registered: true,
		extensions: []string{
			"tsd",
		},
	},
	{
		name:       "application/tlsrpt+gzip",
		format:     "application/x-gzip",
		registered: true,
	},
	{
		name:       "application/tlsrpt+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/tm+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/tnauthlist",
		registered: true,
	},
	{
		name:       "application/token-introspection+jwt",
		format:     "application/jwt",
		registered: true,
	},
	{
		name:       "application/toolbook",
		registered: false,
		extensions: []string{
			"tbk",
		},
	},
	{
		name:       "application/trickle-ice-sdpfrag",
		registered: true,
	},
	{
		name:       "application/trig",
		registered: true,
	},
	{
		name:       "application/ttml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/tve-trigger",
		registered: true,
	},
	{
		name:       "application/tzif",
		registered: true,
	},
	{
		name:       "application/tzif-leap",
		registered: true,
	},
	{
		name:       "application/ulpfec",
		registered: true,
	},
	{
		name:       "application/urc-grpsheet+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/urc-ressheet+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/urc-targetdesc+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/urc-uisocketdesc+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vcard+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vcard+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vda",
		registered: false,
		extensions: []string{
			"vda",
		},
	},
	{
		name:       "application/vemmi",
		registered: true,
	},
	{
		name:       "application/vividence.scriptfile",
		registered: false,
	},
	{
		name:       "application/vnd.1000minds.decision-model+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3M.Post-it-Notes",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp-prose+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp-prose-pc3a+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp-prose-pc3ach+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp-prose-pc3ch+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp-prose-pc8+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp-v2x-local-service-information",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.5gnas",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.GMOP+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.SRVCC-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.access-transfer-events+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.bsf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.gtpc",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.interworking-data",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.lpp",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mc-signalling-ear",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-affiliation-command+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-msgstore-ctrl-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-payload",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-regroup+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-service-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-signalling",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-ue-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcdata-user-profile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-affiliation-command+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-floor-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-location-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-mbms-usage-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-service-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-signed+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-ue-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-ue-init-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcptt-user-profile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-affiliation-command+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-affiliation-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-location-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-mbms-usage-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-service-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-transmission-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-ue-config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mcvideo-user-profile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.mid-call+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.ngap",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.pfcp",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.pic-bw-large",
		registered: true,
		extensions: []string{
			"plb",
		},
	},
	{
		name:       "application/vnd.3gpp.pic-bw-small",
		registered: true,
		extensions: []string{
			"psb",
		},
	},
	{
		name:       "application/vnd.3gpp.pic-bw-var",
		registered: true,
		extensions: []string{
			"pvb",
		},
	},
	{
		name:       "application/vnd.3gpp.s1ap",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.sms",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.sms+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.srvcc-ext+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.state-and-event-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp.ussd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp2.bcmcsinfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp2.sms",
		registered: true,
	},
	{
		name:       "application/vnd.3gpp2.tcap",
		registered: true,
		extensions: []string{
			"tcap",
		},
	},
	{
		name:       "application/vnd.3lightssoftware.imagescal",
		registered: true,
	},
	{
		name:       "application/vnd.HandHeld-Entertainment+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.Kinar",
		registered: true,
	},
	{
		name:       "application/vnd.Mobius.DAF",
		registered: true,
	},
	{
		name:       "application/vnd.Mobius.DIS",
		registered: true,
	},
	{
		name:       "application/vnd.Mobius.MBK",
		registered: true,
	},
	{
		name:       "application/vnd.Mobius.MQY",
		registered: true,
	},
	{
		name:       "application/vnd.Mobius.MSL",
		registered: true,
	},
	{
		name:       "application/vnd.Quark.QuarkXPress",
		registered: true,
	},
	{
		name:       "application/vnd.SimTech-MindMapper",
		registered: true,
	},
	{
		name:       "application/vnd.accpac.simply.aso",
		registered: true,
		extensions: []string{
			"aso",
		},
	},
	{
		name:       "application/vnd.accpac.simply.imp",
		registered: true,
		extensions: []string{
			"imp",
		},
	},
	{
		name:       "application/vnd.acucobol",
		registered: true,
		extensions: []string{
			"acu",
		},
	},
	{
		name:       "application/vnd.acucorp",
		registered: true,
		extensions: []string{
			"atc", "acutc",
		},
	},
	{
		name:       "application/vnd.adobe.air-application-installer-package+zip",
		format:     "application/zip",
		registered: false,
		extensions: []string{
			"air",
		},
	},
	{
		name:       "application/vnd.adobe.flash.movie",
		registered: true,
	},
	{
		name:       "application/vnd.adobe.formscentral.fcdt",
		registered: true,
		extensions: []string{
			"fcdt",
		},
	},
	{
		name:       "application/vnd.adobe.fxp",
		registered: true,
		extensions: []string{
			"fxp", "fxpl",
		},
	},
	{
		name:       "application/vnd.adobe.partial-upload",
		registered: true,
	},
	{
		name:       "application/vnd.adobe.xdp+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xdp",
		},
	},
	{
		name:       "application/vnd.adobe.xfdf",
		registered: false,
		extensions: []string{
			"xfdf",
		},
	},
	{
		name:       "application/vnd.aether.imp",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.afplinedata",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.afplinedata-pagedef",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.cmoca-cmresource",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.foca-charset",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.foca-codedfont",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.foca-codepage",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca-cmtable",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca-formdef",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca-mediummap",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca-objectcontainer",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca-overlay",
		registered: true,
	},
	{
		name:       "application/vnd.afpc.modca-pagesegment",
		registered: true,
	},
	{
		name:       "application/vnd.age",
		registered: true,
	},
	{
		name:       "application/vnd.ah-barcode",
		registered: true,
	},
	{
		name:       "application/vnd.ahead.space",
		registered: true,
		extensions: []string{
			"ahead",
		},
	},
	{
		name:       "application/vnd.airzip.filesecure.azf",
		registered: true,
		extensions: []string{
			"azf",
		},
	},
	{
		name:       "application/vnd.airzip.filesecure.azs",
		registered: true,
		extensions: []string{
			"azs",
		},
	},
	{
		name:       "application/vnd.amadeus+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.amazon.ebook",
		registered: false,
		extensions: []string{
			"azw",
		},
	},
	{
		name:       "application/vnd.amazon.mobi8-ebook",
		registered: true,
	},
	{
		name:       "application/vnd.americandynamics.acc",
		registered: true,
		extensions: []string{
			"acc",
		},
	},
	{
		name:       "application/vnd.amiga.ami",
		registered: true,
		extensions: []string{
			"ami",
		},
	},
	{
		name:       "application/vnd.amundsen.maze+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.android.ota",
		registered: true,
	},
	{
		name:       "application/vnd.android.package-archive",
		registered: false,
		extensions: []string{
			"apk",
		},
	},
	{
		name:       "application/vnd.anki",
		registered: true,
	},
	{
		name:       "application/vnd.anser-web-certificate-issue-initiation",
		registered: true,
		extensions: []string{
			"cii",
		},
	},
	{
		name:       "application/vnd.anser-web-funds-transfer-initiation",
		registered: false,
		extensions: []string{
			"fti",
		},
	},
	{
		name:       "application/vnd.antix.game-component",
		registered: true,
		extensions: []string{
			"atx",
		},
	},
	{
		name:       "application/vnd.apache.arrow.file",
		registered: true,
	},
	{
		name:       "application/vnd.apache.arrow.stream",
		registered: true,
	},
	{
		name:       "application/vnd.apache.thrift.binary",
		registered: true,
	},
	{
		name:       "application/vnd.apache.thrift.compact",
		registered: true,
	},
	{
		name:       "application/vnd.apache.thrift.json",
		registered: true,
	},
	{
		name:       "application/vnd.apexlang",
		registered: true,
	},
	{
		name:       "application/vnd.api+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.aplextor.warrp+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.apothekende.reservation+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.apple.installer+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mpkg",
		},
	},
	{
		name:       "application/vnd.apple.keynote",
		registered: true,
	},
	{
		name:       "application/vnd.apple.mpegurl",
		registered: true,
		extensions: []string{
			"m3u8",
		},
	},
	{
		name:       "application/vnd.apple.numbers",
		registered: true,
	},
	{
		name:       "application/vnd.apple.pages",
		registered: true,
	},
	{
		name:       "application/vnd.arastra.swi",
		registered: true,
		extensions: []string{
			"swi",
		},
	},
	{
		name:       "application/vnd.aristanetworks.swi",
		registered: true,
		extensions: []string{
			"swi",
		},
	},
	{
		name:       "application/vnd.artisan+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.artsquare",
		registered: true,
	},
	{
		name:       "application/vnd.astraea-software.iota",
		registered: true,
		extensions: []string{
			"iota",
		},
	},
	{
		name:       "application/vnd.audiograph",
		registered: true,
		extensions: []string{
			"aep",
		},
	},
	{
		name:       "application/vnd.autopackage",
		registered: true,
	},
	{
		name:       "application/vnd.avalon+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.avistar+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.balsamiq.bmml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.balsamiq.bmpr",
		registered: true,
	},
	{
		name:       "application/vnd.banana-accounting",
		registered: true,
	},
	{
		name:       "application/vnd.bbf.usp.error",
		registered: true,
	},
	{
		name:       "application/vnd.bbf.usp.msg",
		registered: true,
	},
	{
		name:       "application/vnd.bbf.usp.msg+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.bekitzur-stech+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.belightsoft.lhzd+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.belightsoft.lhzl+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.bint.med-content",
		registered: true,
	},
	{
		name:       "application/vnd.biopax.rdf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.blink-idb-value-wrapper",
		registered: true,
	},
	{
		name:       "application/vnd.blueice.multipass",
		registered: true,
		extensions: []string{
			"mpm",
		},
	},
	{
		name:       "application/vnd.bluetooth.ep.oob",
		registered: true,
	},
	{
		name:       "application/vnd.bluetooth.le.oob",
		registered: true,
	},
	{
		name:       "application/vnd.bmi",
		registered: true,
		extensions: []string{
			"bmi",
		},
	},
	{
		name:       "application/vnd.bpf",
		registered: true,
	},
	{
		name:       "application/vnd.bpf3",
		registered: true,
	},
	{
		name:       "application/vnd.businessobjects",
		registered: true,
		extensions: []string{
			"rep",
		},
	},
	{
		name:       "application/vnd.byu.uapi+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.cab-jscript",
		registered: true,
	},
	{
		name:       "application/vnd.canon-cpdl",
		registered: true,
	},
	{
		name:       "application/vnd.canon-lips",
		registered: true,
	},
	{
		name:       "application/vnd.capasystems-pg+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.cendio.thinlinc.clientconf",
		registered: true,
	},
	{
		name:       "application/vnd.century-systems.tcp_stream",
		registered: true,
	},
	{
		name:       "application/vnd.chemdraw+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"cdxml",
		},
	},
	{
		name:       "application/vnd.chess-pgn",
		registered: true,
	},
	{
		name:       "application/vnd.chipnuts.karaoke-mmd",
		registered: true,
		extensions: []string{
			"mmd",
		},
	},
	{
		name:       "application/vnd.ciedi",
		registered: true,
	},
	{
		name:       "application/vnd.cinderella",
		registered: true,
		extensions: []string{
			"cdy",
		},
	},
	{
		name:       "application/vnd.cirpack.isdn-ext",
		registered: true,
	},
	{
		name:       "application/vnd.citationstyles.style+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.claymore",
		registered: true,
		extensions: []string{
			"cla",
		},
	},
	{
		name:       "application/vnd.cloanto.rp9",
		registered: true,
		extensions: []string{
			"rp9",
		},
	},
	{
		name:       "application/vnd.clonk.c4group",
		registered: true,
		extensions: []string{
			"c4g", "c4d", "c4f", "c4p", "c4u",
		},
	},
	{
		name:       "application/vnd.cluetrust.cartomobile-config",
		registered: true,
		extensions: []string{
			"c11amc",
		},
	},
	{
		name:       "application/vnd.cluetrust.cartomobile-config-pkg",
		registered: true,
		extensions: []string{
			"c11amz",
		},
	},
	{
		name:       "application/vnd.cncf.helm.chart.content.v1.tar+gzip",
		format:     "application/x-gzip",
		registered: true,
	},
	{
		name:       "application/vnd.cncf.helm.chart.provenance.v1.prov",
		registered: true,
	},
	{
		name:       "application/vnd.coffeescript",
		registered: true,
	},
	{
		name:       "application/vnd.collabio.xodocuments.document",
		registered: true,
	},
	{
		name:       "application/vnd.collabio.xodocuments.document-template",
		registered: true,
	},
	{
		name:       "application/vnd.collabio.xodocuments.presentation",
		registered: true,
	},
	{
		name:       "application/vnd.collabio.xodocuments.presentation-template",
		registered: true,
	},
	{
		name:       "application/vnd.collabio.xodocuments.spreadsheet",
		registered: true,
	},
	{
		name:       "application/vnd.collabio.xodocuments.spreadsheet-template",
		registered: true,
	},
	{
		name:       "application/vnd.collection+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.collection.doc+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.collection.next+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.comicbook+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.comicbook-rar",
		registered: true,
	},
	{
		name:       "application/vnd.commerce-battelle",
		registered: true,
	},
	{
		name:       "application/vnd.commonspace",
		registered: true,
		extensions: []string{
			"csp",
		},
	},
	{
		name:       "application/vnd.comsocaller",
		registered: false,
	},
	{
		name:       "application/vnd.contact.cmsg",
		registered: true,
		extensions: []string{
			"cdbcmsg",
		},
	},
	{
		name:       "application/vnd.coreos.ignition+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.cosmocaller",
		registered: true,
		extensions: []string{
			"cmc",
		},
	},
	{
		name:       "application/vnd.crick.clicker",
		registered: true,
		extensions: []string{
			"clkx",
		},
	},
	{
		name:       "application/vnd.crick.clicker.keyboard",
		registered: true,
		extensions: []string{
			"clkk",
		},
	},
	{
		name:       "application/vnd.crick.clicker.palette",
		registered: true,
		extensions: []string{
			"clkp",
		},
	},
	{
		name:       "application/vnd.crick.clicker.template",
		registered: true,
		extensions: []string{
			"clkt",
		},
	},
	{
		name:       "application/vnd.crick.clicker.wordbank",
		registered: true,
		extensions: []string{
			"clkw",
		},
	},
	{
		name:       "application/vnd.criticaltools.wbs+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"wbs",
		},
	},
	{
		name:       "application/vnd.cryptii.pipe+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.crypto-shade-file",
		registered: true,
	},
	{
		name:       "application/vnd.cryptomator.encrypted",
		registered: true,
	},
	{
		name:       "application/vnd.cryptomator.vault",
		registered: true,
	},
	{
		name:       "application/vnd.ctc-posml",
		registered: true,
		extensions: []string{
			"pml",
		},
	},
	{
		name:       "application/vnd.ctct.ws+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.cups-pdf",
		registered: true,
	},
	{
		name:       "application/vnd.cups-postscript",
		registered: true,
	},
	{
		name:       "application/vnd.cups-ppd",
		registered: true,
		extensions: []string{
			"ppd",
		},
	},
	{
		name:       "application/vnd.cups-raster",
		registered: true,
	},
	{
		name:       "application/vnd.cups-raw",
		registered: true,
	},
	{
		name:       "application/vnd.curl",
		registered: true,
	},
	{
		name:       "application/vnd.curl.car",
		registered: false,
		extensions: []string{
			"car",
		},
	},
	{
		name:       "application/vnd.curl.pcurl",
		registered: false,
		extensions: []string{
			"pcurl",
		},
	},
	{
		name:       "application/vnd.cyan.dean.root+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.cybank",
		registered: true,
	},
	{
		name:       "application/vnd.cyclonedx+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.cyclonedx+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.d2l.coursepackage1p0+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.d3m-dataset",
		registered: true,
	},
	{
		name:       "application/vnd.d3m-problem",
		registered: true,
	},
	{
		name:       "application/vnd.dart",
		registered: true,
		extensions: []string{
			"dart",
		},
	},
	{
		name:       "application/vnd.data-vision.rdz",
		registered: true,
		extensions: []string{
			"rdz",
		},
	},
	{
		name:       "application/vnd.datalog",
		registered: true,
	},
	{
		name:       "application/vnd.datapackage+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.dataresource+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.dbf",
		registered: true,
	},
	{
		name:       "application/vnd.debian.binary-package",
		registered: true,
	},
	{
		name:       "application/vnd.dece.data",
		registered: true,
		extensions: []string{
			"uvf", "uvvf", "uvd", "uvvd",
		},
	},
	{
		name:       "application/vnd.dece.ttml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"uvt", "uvvt",
		},
	},
	{
		name:       "application/vnd.dece.unspecified",
		registered: true,
		extensions: []string{
			"uvx", "uvvx",
		},
	},
	{
		name:       "application/vnd.dece.zip",
		registered: true,
		extensions: []string{
			"uvz", "uvvz",
		},
	},
	{
		name:       "application/vnd.denovo.fcselayout-link",
		registered: true,
		extensions: []string{
			"fe_launch",
		},
	},
	{
		name:       "application/vnd.desmume.movie",
		registered: true,
	},
	{
		name:       "application/vnd.dir-bi.plate-dl-nosuffix",
		registered: true,
	},
	{
		name:       "application/vnd.dm.delegation+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dna",
		registered: true,
		extensions: []string{
			"dna",
		},
	},
	{
		name:       "application/vnd.document+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.dolby.mlp",
		registered: false,
		extensions: []string{
			"mlp",
		},
	},
	{
		name:       "application/vnd.dolby.mobile.1",
		registered: true,
	},
	{
		name:       "application/vnd.dolby.mobile.2",
		registered: true,
	},
	{
		name:       "application/vnd.doremir.scorecloud-binary-document",
		registered: true,
	},
	{
		name:       "application/vnd.dpgraph",
		registered: true,
		extensions: []string{
			"dpg",
		},
	},
	{
		name:       "application/vnd.dreamfactory",
		registered: true,
		extensions: []string{
			"dfac",
		},
	},
	{
		name:       "application/vnd.drive+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ds-keypoint",
		registered: false,
		extensions: []string{
			"kpxx",
		},
	},
	{
		name:       "application/vnd.dtg.local",
		registered: true,
	},
	{
		name:       "application/vnd.dtg.local.flash",
		registered: true,
	},
	{
		name:       "application/vnd.dtg.local.html",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.ait",
		registered: true,
		extensions: []string{
			"ait",
		},
	},
	{
		name:       "application/vnd.dvb.dvbisl+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.dvbj",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.esgcontainer",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.ipdcdftnotifaccess",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.ipdcesgaccess",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.ipdcesgaccess2",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.ipdcesgpdd",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.ipdcroaming",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.iptv.alfec-base",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.iptv.alfec-enhancement",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-aggregate-root+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-container+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-generic+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-ia-msglist+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-ia-registration-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-ia-registration-response+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.notif-init+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.pfr",
		registered: true,
	},
	{
		name:       "application/vnd.dvb.service",
		registered: true,
		extensions: []string{
			"svc",
		},
	},
	{
		name:       "application/vnd.dxr",
		registered: true,
	},
	{
		name:       "application/vnd.dynageo",
		registered: true,
		extensions: []string{
			"geo",
		},
	},
	{
		name:       "application/vnd.dzr",
		registered: true,
	},
	{
		name:       "application/vnd.easykaraoke.cdgdownload",
		registered: true,
	},
	{
		name:       "application/vnd.ecdis-update",
		registered: true,
	},
	{
		name:       "application/vnd.ecip.rlp",
		registered: true,
	},
	{
		name:       "application/vnd.eclipse.ditto+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ecowin.chart",
		registered: true,
		extensions: []string{
			"mag",
		},
	},
	{
		name:       "application/vnd.ecowin.filerequest",
		registered: true,
	},
	{
		name:       "application/vnd.ecowin.fileupdate",
		registered: true,
	},
	{
		name:       "application/vnd.ecowin.series",
		registered: true,
	},
	{
		name:       "application/vnd.ecowin.seriesrequest",
		registered: true,
	},
	{
		name:       "application/vnd.ecowin.seriesupdate",
		registered: true,
	},
	{
		name:       "application/vnd.efi.img",
		registered: true,
	},
	{
		name:       "application/vnd.efi.iso",
		registered: true,
	},
	{
		name:       "application/vnd.emclient.accessrequest+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.enliven",
		registered: true,
		extensions: []string{
			"nml",
		},
	},
	{
		name:       "application/vnd.enphase.envoy",
		registered: true,
	},
	{
		name:       "application/vnd.eprints.data+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.epson.esf",
		registered: true,
		extensions: []string{
			"esf",
		},
	},
	{
		name:       "application/vnd.epson.msf",
		registered: true,
		extensions: []string{
			"msf",
		},
	},
	{
		name:       "application/vnd.epson.quickanime",
		registered: true,
		extensions: []string{
			"qam",
		},
	},
	{
		name:       "application/vnd.epson.salt",
		registered: true,
		extensions: []string{
			"slt",
		},
	},
	{
		name:       "application/vnd.epson.ssf",
		registered: true,
		extensions: []string{
			"ssf",
		},
	},
	{
		name:       "application/vnd.ericsson.quickcall",
		registered: true,
	},
	{
		name:       "application/vnd.espass-espass+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.eszigno3+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"es3", "et3",
		},
	},
	{
		name:       "application/vnd.etsi.aoc+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.asic-e+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.asic-s+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.cug+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvcommand+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvdiscovery+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvprofile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvsad-bc+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvsad-cod+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvsad-npvr+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvservice+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvsync+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.iptvueprofile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.mcid+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.mheg5",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.overload-control-policy-dataset+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.pstn+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.sci+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.simservs+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.timestamp-token",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.tsl+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.etsi.tsl.der",
		registered: true,
	},
	{
		name:       "application/vnd.eu.kasparian.car+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.eudora.data",
		registered: true,
	},
	{
		name:       "application/vnd.evolv.ecig.profile",
		registered: true,
	},
	{
		name:       "application/vnd.evolv.ecig.settings",
		registered: true,
	},
	{
		name:       "application/vnd.evolv.ecig.theme",
		registered: true,
	},
	{
		name:       "application/vnd.exstream-empower+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.exstream-package",
		registered: true,
	},
	{
		name:       "application/vnd.ezpix-album",
		registered: true,
		extensions: []string{
			"ez2",
		},
	},
	{
		name:       "application/vnd.ezpix-package",
		registered: true,
		extensions: []string{
			"ez3",
		},
	},
	{
		name:       "application/vnd.f-secure.mobile",
		registered: true,
	},
	{
		name:       "application/vnd.familysearch.gedcom+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.fastcopy-disk-image",
		registered: true,
	},
	{
		name:       "application/vnd.fdf",
		registered: false,
		extensions: []string{
			"fdf",
		},
	},
	{
		name:       "application/vnd.fdsn.mseed",
		registered: true,
		extensions: []string{
			"mseed",
		},
	},
	{
		name:       "application/vnd.fdsn.seed",
		registered: true,
		extensions: []string{
			"seed", "dataless",
		},
	},
	{
		name:       "application/vnd.ffsns",
		registered: true,
	},
	{
		name:       "application/vnd.ficlab.flb+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.filmit.zfc",
		registered: true,
	},
	{
		name:       "application/vnd.fints",
		registered: true,
	},
	{
		name:       "application/vnd.firemonkeys.cloudcell",
		registered: true,
	},
	{
		name:       "application/vnd.flographit",
		registered: false,
		extensions: []string{
			"gph",
		},
	},
	{
		name:       "application/vnd.fluxtime.clip",
		registered: true,
		extensions: []string{
			"ftc",
		},
	},
	{
		name:       "application/vnd.font-fontforge-sfd",
		registered: true,
	},
	{
		name:       "application/vnd.framemaker",
		registered: true,
		extensions: []string{
			"fm", "frame", "maker", "book",
		},
	},
	{
		name:       "application/vnd.frogans.fnc",
		registered: true,
		extensions: []string{
			"fnc",
		},
	},
	{
		name:       "application/vnd.frogans.ltf",
		registered: true,
		extensions: []string{
			"ltf",
		},
	},
	{
		name:       "application/vnd.fsc.weblaunch",
		registered: true,
		extensions: []string{
			"fsc",
		},
	},
	{
		name:       "application/vnd.fujifilm.fb.docuworks",
		registered: true,
	},
	{
		name:       "application/vnd.fujifilm.fb.docuworks.binder",
		registered: true,
	},
	{
		name:       "application/vnd.fujifilm.fb.docuworks.container",
		registered: true,
	},
	{
		name:       "application/vnd.fujifilm.fb.jfi+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.fujitsu.oasys",
		registered: true,
		extensions: []string{
			"oas",
		},
	},
	{
		name:       "application/vnd.fujitsu.oasys2",
		registered: true,
		extensions: []string{
			"oa2",
		},
	},
	{
		name:       "application/vnd.fujitsu.oasys3",
		registered: true,
		extensions: []string{
			"oa3",
		},
	},
	{
		name:       "application/vnd.fujitsu.oasysgp",
		registered: true,
		extensions: []string{
			"fg5",
		},
	},
	{
		name:       "application/vnd.fujitsu.oasysprs",
		registered: true,
		extensions: []string{
			"bh2",
		},
	},
	{
		name:       "application/vnd.fujixerox.HBPL",
		registered: true,
	},
	{
		name:       "application/vnd.fujixerox.art-ex",
		registered: false,
	},
	{
		name:       "application/vnd.fujixerox.art4",
		registered: false,
	},
	{
		name:       "application/vnd.fujixerox.ddd",
		registered: true,
		extensions: []string{
			"ddd",
		},
	},
	{
		name:       "application/vnd.fujixerox.docuworks",
		registered: true,
		extensions: []string{
			"xdw",
		},
	},
	{
		name:       "application/vnd.fujixerox.docuworks.binder",
		registered: true,
		extensions: []string{
			"xbd",
		},
	},
	{
		name:       "application/vnd.fujixerox.docuworks.container",
		registered: true,
	},
	{
		name:       "application/vnd.fut-misnet",
		registered: true,
	},
	{
		name:       "application/vnd.futoin+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/vnd.futoin+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.fuzzysheet",
		registered: true,
		extensions: []string{
			"fzs",
		},
	},
	{
		name:       "application/vnd.genomatix.tuxedo",
		registered: true,
		extensions: []string{
			"txd",
		},
	},
	{
		name:       "application/vnd.genozip",
		registered: true,
	},
	{
		name:       "application/vnd.gentics.grd+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.catmetadata+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.ebuild",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.eclass",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.gpkg",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.manifest",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.pkgmetadata+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.gentoo.xpak",
		registered: true,
	},
	{
		name:       "application/vnd.geo+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.geocube+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.geogebra.file",
		registered: true,
		extensions: []string{
			"ggb",
		},
	},
	{
		name:       "application/vnd.geogebra.slides",
		registered: true,
	},
	{
		name:       "application/vnd.geogebra.tool",
		registered: true,
		extensions: []string{
			"ggt",
		},
	},
	{
		name:       "application/vnd.geometry-explorer",
		registered: true,
		extensions: []string{
			"gex", "gre",
		},
	},
	{
		name:       "application/vnd.geonext",
		registered: true,
		extensions: []string{
			"gxt",
		},
	},
	{
		name:       "application/vnd.geoplan",
		registered: true,
		extensions: []string{
			"g2w",
		},
	},
	{
		name:       "application/vnd.geospace",
		registered: true,
		extensions: []string{
			"g3w",
		},
	},
	{
		name:       "application/vnd.gerber",
		registered: true,
	},
	{
		name:       "application/vnd.globalplatform.card-content-mgt",
		registered: true,
	},
	{
		name:       "application/vnd.globalplatform.card-content-mgt-response",
		registered: true,
	},
	{
		name:       "application/vnd.gmx",
		registered: true,
		extensions: []string{
			"gmx",
		},
	},
	{
		name:       "application/vnd.gnu.taler.exchange+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.gnu.taler.merchant+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.google-earth.kml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"kml",
		},
	},
	{
		name:       "application/vnd.google-earth.kmz",
		registered: true,
		extensions: []string{
			"kmz",
		},
	},
	{
		name:       "application/vnd.gov.sk.e-form+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.gov.sk.e-form+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.gov.sk.xmldatacontainer+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.grafeq",
		registered: true,
		extensions: []string{
			"gqf", "gqs",
		},
	},
	{
		name:       "application/vnd.gridmp",
		registered: true,
	},
	{
		name:       "application/vnd.groove-account",
		registered: true,
		extensions: []string{
			"gac",
		},
	},
	{
		name:       "application/vnd.groove-help",
		registered: true,
		extensions: []string{
			"ghf",
		},
	},
	{
		name:       "application/vnd.groove-identity-message",
		registered: true,
		extensions: []string{
			"gim",
		},
	},
	{
		name:       "application/vnd.groove-injector",
		registered: true,
		extensions: []string{
			"grv",
		},
	},
	{
		name:       "application/vnd.groove-tool-message",
		registered: true,
		extensions: []string{
			"gtm",
		},
	},
	{
		name:       "application/vnd.groove-tool-template",
		registered: true,
		extensions: []string{
			"tpl",
		},
	},
	{
		name:       "application/vnd.groove-vcard",
		registered: true,
		extensions: []string{
			"vcg",
		},
	},
	{
		name:       "application/vnd.hal+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.hal+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"hal",
		},
	},
	{
		name:       "application/vnd.hbci",
		registered: true,
		extensions: []string{
			"hbci",
		},
	},
	{
		name:       "application/vnd.hc+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.hcl-bireports",
		registered: true,
	},
	{
		name:       "application/vnd.hdt",
		registered: true,
	},
	{
		name:       "application/vnd.heroku+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.hhe.lesson-player",
		registered: true,
		extensions: []string{
			"les",
		},
	},
	{
		name:       "application/vnd.hp-PCLXL",
		registered: true,
	},
	{
		name:       "application/vnd.hp-hpgl",
		registered: false,
		extensions: []string{
			"hgl", "hpg", "hpgl",
		},
	},
	{
		name:       "application/vnd.hp-hpid",
		registered: true,
		extensions: []string{
			"hpid",
		},
	},
	{
		name:       "application/vnd.hp-hps",
		registered: true,
		extensions: []string{
			"hps",
		},
	},
	{
		name:       "application/vnd.hp-jlyt",
		registered: true,
		extensions: []string{
			"jlt",
		},
	},
	{
		name:       "application/vnd.hp-pcl",
		registered: false,
		extensions: []string{
			"pcl",
		},
	},
	{
		name:       "application/vnd.httphone",
		registered: true,
	},
	{
		name:       "application/vnd.hydrostatix.sof-data",
		registered: true,
		extensions: []string{
			"sfd-hdstx",
		},
	},
	{
		name:       "application/vnd.hyper+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.hyper-item+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.hyperdrive+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.hzn-3d-crossword",
		registered: true,
		extensions: []string{
			"x3d",
		},
	},
	{
		name:       "application/vnd.ibm.MiniPay",
		registered: true,
	},
	{
		name:       "application/vnd.ibm.afplinedata",
		registered: true,
	},
	{
		name:       "application/vnd.ibm.electronic-media",
		registered: true,
	},
	{
		name:       "application/vnd.ibm.modcap",
		registered: true,
		extensions: []string{
			"afp", "listafp", "list3820",
		},
	},
	{
		name:       "application/vnd.ibm.rights-management",
		registered: true,
		extensions: []string{
			"irm",
		},
	},
	{
		name:       "application/vnd.ibm.secure-container",
		registered: true,
		extensions: []string{
			"sc",
		},
	},
	{
		name:       "application/vnd.iccprofile",
		registered: true,
		extensions: []string{
			"icc", "icm",
		},
	},
	{
		name:       "application/vnd.ieee.1905",
		registered: true,
	},
	{
		name:       "application/vnd.igloader",
		registered: true,
		extensions: []string{
			"igl",
		},
	},
	{
		name:       "application/vnd.imagemeter.folder+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.imagemeter.image+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.immervision-ivp",
		registered: true,
		extensions: []string{
			"ivp",
		},
	},
	{
		name:       "application/vnd.immervision-ivu",
		registered: true,
		extensions: []string{
			"ivu",
		},
	},
	{
		name:       "application/vnd.ims.imsccv1p1",
		registered: true,
	},
	{
		name:       "application/vnd.ims.imsccv1p2",
		registered: true,
	},
	{
		name:       "application/vnd.ims.imsccv1p3",
		registered: true,
	},
	{
		name:       "application/vnd.ims.lis.v2.result+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ims.lti.v2.toolconsumerprofile+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ims.lti.v2.toolproxy+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ims.lti.v2.toolproxy.id+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ims.lti.v2.toolsettings+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ims.lti.v2.toolsettings.simple+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.informedcontrol.rms+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.informix-visionary",
		registered: true,
	},
	{
		name:       "application/vnd.infotech.project",
		registered: true,
	},
	{
		name:       "application/vnd.infotech.project+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.innopath.wamp.notification",
		registered: true,
	},
	{
		name:       "application/vnd.insors.igm",
		registered: true,
		extensions: []string{
			"igm",
		},
	},
	{
		name:       "application/vnd.intercon.formnet",
		registered: true,
		extensions: []string{
			"xpw", "xpx",
		},
	},
	{
		name:       "application/vnd.intergeo",
		registered: true,
		extensions: []string{
			"i2g",
		},
	},
	{
		name:       "application/vnd.intertrust.digibox",
		registered: true,
	},
	{
		name:       "application/vnd.intertrust.nncp",
		registered: true,
	},
	{
		name:       "application/vnd.intu.qbo",
		registered: true,
		extensions: []string{
			"qbo",
		},
	},
	{
		name:       "application/vnd.intu.qfx",
		registered: true,
		extensions: []string{
			"qfx",
		},
	},
	{
		name:       "application/vnd.ipld.car",
		registered: true,
	},
	{
		name:       "application/vnd.ipld.dag-cbor",
		registered: true,
	},
	{
		name:       "application/vnd.ipld.dag-json",
		registered: true,
	},
	{
		name:       "application/vnd.ipld.raw",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.catalogitem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.conceptitem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.knowledgeitem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.newsitem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.newsmessage+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.packageitem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.iptc.g2.planningitem+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.ipunplugged.rcprofile",
		registered: true,
		extensions: []string{
			"rcprofile",
		},
	},
	{
		name:       "application/vnd.irepository.package+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"irp",
		},
	},
	{
		name:       "application/vnd.is-xpr",
		registered: true,
		extensions: []string{
			"xpr",
		},
	},
	{
		name:       "application/vnd.isac.fcs",
		registered: true,
		extensions: []string{
			"fcs",
		},
	},
	{
		name:       "application/vnd.iso11783-10+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.jam",
		registered: true,
		extensions: []string{
			"jam",
		},
	},
	{
		name:       "application/vnd.japannet-directory-service",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-jpnstore-wakeup",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-payment-wakeup",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-registration",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-registration-wakeup",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-setstore-wakeup",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-verification",
		registered: true,
	},
	{
		name:       "application/vnd.japannet-verification-wakeup",
		registered: true,
	},
	{
		name:       "application/vnd.jcp.javame.midlet-rms",
		registered: true,
		extensions: []string{
			"rms",
		},
	},
	{
		name:       "application/vnd.jisp",
		registered: true,
		extensions: []string{
			"jisp",
		},
	},
	{
		name:       "application/vnd.joost.joda-archive",
		registered: true,
		extensions: []string{
			"joda",
		},
	},
	{
		name:       "application/vnd.jsk.isdn-ngn",
		registered: true,
	},
	{
		name:       "application/vnd.kahootz",
		registered: true,
		extensions: []string{
			"ktz", "ktr",
		},
	},
	{
		name:       "application/vnd.kde.karbon",
		registered: true,
		extensions: []string{
			"karbon",
		},
	},
	{
		name:       "application/vnd.kde.kchart",
		registered: true,
		extensions: []string{
			"chrt",
		},
	},
	{
		name:       "application/vnd.kde.kformula",
		registered: true,
		extensions: []string{
			"kfo",
		},
	},
	{
		name:       "application/vnd.kde.kivio",
		registered: true,
		extensions: []string{
			"flw",
		},
	},
	{
		name:       "application/vnd.kde.kontour",
		registered: true,
		extensions: []string{
			"kon",
		},
	},
	{
		name:       "application/vnd.kde.kpresenter",
		registered: true,
		extensions: []string{
			"kpr", "kpt",
		},
	},
	{
		name:       "application/vnd.kde.kspread",
		registered: true,
		extensions: []string{
			"ksp",
		},
	},
	{
		name:       "application/vnd.kde.kword",
		registered: true,
		extensions: []string{
			"kwd", "kwt",
		},
	},
	{
		name:       "application/vnd.kenameaapp",
		registered: true,
		extensions: []string{
			"htke",
		},
	},
	{
		name:       "application/vnd.kidspiration",
		registered: true,
		extensions: []string{
			"kia",
		},
	},
	{
		name:       "application/vnd.koan",
		registered: true,
		extensions: []string{
			"skp", "skd", "skt", "skm",
		},
	},
	{
		name:       "application/vnd.kodak-descriptor",
		registered: true,
		extensions: []string{
			"sse",
		},
	},
	{
		name:       "application/vnd.las",
		registered: true,
	},
	{
		name:       "application/vnd.las.las+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.las.las+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"lasxml",
		},
	},
	{
		name:       "application/vnd.laszip",
		registered: true,
	},
	{
		name:       "application/vnd.leap+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.liberty-request+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.llamagraphics.life-balance.desktop",
		registered: true,
		extensions: []string{
			"lbd",
		},
	},
	{
		name:       "application/vnd.llamagraphics.life-balance.exchange+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"lbe",
		},
	},
	{
		name:       "application/vnd.logipipe.circuit+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.loom",
		registered: true,
	},
	{
		name:       "application/vnd.lotus-1-2-3",
		registered: true,
		extensions: []string{
			"123",
		},
	},
	{
		name:       "application/vnd.lotus-approach",
		registered: true,
		extensions: []string{
			"apr",
		},
	},
	{
		name:       "application/vnd.lotus-freelance",
		registered: true,
		extensions: []string{
			"pre",
		},
	},
	{
		name:       "application/vnd.lotus-notes",
		registered: true,
		extensions: []string{
			"nsf",
		},
	},
	{
		name:       "application/vnd.lotus-organizer",
		registered: true,
		extensions: []string{
			"org",
		},
	},
	{
		name:       "application/vnd.lotus-screencam",
		registered: true,
		extensions: []string{
			"scm",
		},
	},
	{
		name:       "application/vnd.lotus-wordpro",
		registered: true,
		extensions: []string{
			"lwp",
		},
	},
	{
		name:       "application/vnd.macports.portpkg",
		registered: true,
		extensions: []string{
			"portpkg",
		},
	},
	{
		name:       "application/vnd.mapbox-vector-tile",
		registered: true,
	},
	{
		name:       "application/vnd.marlin.drm.actiontoken+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.marlin.drm.conftoken+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.marlin.drm.license+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.marlin.drm.mdcf",
		registered: true,
	},
	{
		name:       "application/vnd.mason+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.maxar.archive.3tz+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "application/vnd.maxmind.maxmind-db",
		registered: true,
	},
	{
		name:       "application/vnd.mcd",
		registered: true,
		extensions: []string{
			"mcd",
		},
	},
	{
		name:       "application/vnd.medcalcdata",
		registered: true,
		extensions: []string{
			"mc1",
		},
	},
	{
		name:       "application/vnd.mediastation.cdkey",
		registered: true,
		extensions: []string{
			"cdkey",
		},
	},
	{
		name:       "application/vnd.meridian-slingshot",
		registered: true,
	},
	{
		name:       "application/vnd.mfer",
		registered: false,
		extensions: []string{
			"mwf",
		},
	},
	{
		name:       "application/vnd.mfmp",
		registered: true,
		extensions: []string{
			"mfm",
		},
	},
	{
		name:       "application/vnd.micro+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.micrografx.flo",
		registered: true,
		extensions: []string{
			"flo",
		},
	},
	{
		name:       "application/vnd.micrografx.igx",
		registered: true,
		extensions: []string{
			"igx",
		},
	},
	{
		name:       "application/vnd.microsoft.portable-executable",
		registered: true,
	},
	{
		name:       "application/vnd.microsoft.windows.thumbnail-cache",
		registered: true,
	},
	{
		name:       "application/vnd.miele+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.mif",
		registered: true,
		extensions: []string{
			"mif",
		},
	},
	{
		name:       "application/vnd.minisoft-hp3000-save",
		registered: true,
	},
	{
		name:       "application/vnd.mitsubishi.misty-guard.trustweb",
		registered: true,
	},
	{
		name:       "application/vnd.mobius.plc",
		registered: false,
		extensions: []string{
			"plc",
		},
	},
	{
		name:       "application/vnd.mobius.txf",
		registered: false,
		extensions: []string{
			"txf",
		},
	},
	{
		name:       "application/vnd.mophun.application",
		registered: true,
		extensions: []string{
			"mpn",
		},
	},
	{
		name:       "application/vnd.mophun.certificate",
		registered: true,
		extensions: []string{
			"mpc",
		},
	},
	{
		name:       "application/vnd.motorola.flexsuite",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.flexsuite.adsi",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.flexsuite.fis",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.flexsuite.gotap",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.flexsuite.kmr",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.flexsuite.ttc",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.flexsuite.wem",
		registered: true,
	},
	{
		name:       "application/vnd.motorola.iprm",
		registered: true,
	},
	{
		name:       "application/vnd.mozilla.xul+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xul",
		},
	},
	{
		name:       "application/vnd.ms-3mfdocument",
		registered: true,
	},
	{
		name:       "application/vnd.ms-PrintDeviceCapabilities+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.ms-PrintSchemaTicket+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.ms-artgalry",
		registered: true,
		extensions: []string{
			"cil",
		},
	},
	{
		name:       "application/vnd.ms-asf",
		registered: true,
	},
	{
		name:       "application/vnd.ms-cab-compressed",
		registered: true,
		extensions: []string{
			"cab",
		},
	},
	{
		name:       "application/vnd.ms-color.iccprofile",
		registered: false,
	},
	{
		name:       "application/vnd.ms-excel",
		registered: true,
		extensions: []string{
			"xls", "xlm", "xla", "xlc", "xlt", "xlb", "xll", "xlw",
		},
	},
	{
		name:       "application/vnd.ms-excel.addin.macroEnabled.12",
		registered: true,
		extensions: []string{
			"xlam",
		},
	},
	{
		name:       "application/vnd.ms-excel.sheet.binary.macroEnabled.12",
		registered: true,
		extensions: []string{
			"xlsb",
		},
	},
	{
		name:       "application/vnd.ms-excel.sheet.macroEnabled.12",
		registered: true,
		extensions: []string{
			"xlsm",
		},
	},
	{
		name:       "application/vnd.ms-excel.template.macroenabled.12",
		registered: false,
		extensions: []string{
			"xltm",
		},
	},
	{
		name:       "application/vnd.ms-fontobject",
		registered: true,
		extensions: []string{
			"eot",
		},
	},
	{
		name:       "application/vnd.ms-htmlhelp",
		registered: true,
		extensions: []string{
			"chm",
		},
	},
	{
		name:       "application/vnd.ms-ims",
		registered: true,
		extensions: []string{
			"ims",
		},
	},
	{
		name:       "application/vnd.ms-lrm",
		registered: true,
		extensions: []string{
			"lrm",
		},
	},
	{
		name:       "application/vnd.ms-office.activeX+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.ms-officetheme",
		registered: true,
		extensions: []string{
			"thmx",
		},
	},
	{
		name:       "application/vnd.ms-opentype",
		registered: false,
	},
	{
		name:       "application/vnd.ms-outlook",
		registered: false,
		extensions: []string{
			"msg",
		},
	},
	{
		name:       "application/vnd.ms-package.obfuscated-opentype",
		registered: false,
	},
	{
		name:       "application/vnd.ms-pki.certstore",
		registered: false,
		extensions: []string{
			"sst",
		},
	},
	{
		name:       "application/vnd.ms-pki.pko",
		registered: false,
		extensions: []string{
			"pko",
		},
	},
	{
		name:       "application/vnd.ms-pki.seccat",
		registered: false,
		extensions: []string{
			"cat",
		},
	},
	{
		name:       "application/vnd.ms-pki.stl",
		registered: false,
		extensions: []string{
			"stl",
		},
	},
	{
		name:       "application/vnd.ms-pkicertstore",
		registered: false,
		extensions: []string{
			"sst",
		},
	},
	{
		name:       "application/vnd.ms-pkiseccat",
		registered: false,
		extensions: []string{
			"cat",
		},
	},
	{
		name:       "application/vnd.ms-pkistl",
		registered: false,
		extensions: []string{
			"stl",
		},
	},
	{
		name:       "application/vnd.ms-playready.initiator+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.ms-powerpoint",
		registered: true,
		extensions: []string{
			"ppt", "pps", "pot", "ppa", "pwz",
		},
	},
	{
		name:       "application/vnd.ms-powerpoint.addin.macroEnabled.12",
		registered: true,
		extensions: []string{
			"ppam",
		},
	},
	{
		name:       "application/vnd.ms-powerpoint.presentation.macroEnabled.12",
		registered: true,
		extensions: []string{
			"pptm", "potm",
		},
	},
	{
		name:       "application/vnd.ms-powerpoint.slide.macroEnabled.12",
		registered: true,
		extensions: []string{
			"sldm",
		},
	},
	{
		name:       "application/vnd.ms-powerpoint.slideshow.macroenabled.12",
		registered: false,
		extensions: []string{
			"ppsm",
		},
	},
	{
		name:       "application/vnd.ms-powerpoint.template.macroEnabled.12",
		registered: true,
		extensions: []string{
			"potm",
		},
	},
	{
		name:       "application/vnd.ms-printing.printticket+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.ms-project",
		registered: true,
		extensions: []string{
			"mpp", "mpt",
		},
	},
	{
		name:       "application/vnd.ms-tnef",
		registered: true,
	},
	{
		name:       "application/vnd.ms-windows.devicepairing",
		registered: true,
	},
	{
		name:       "application/vnd.ms-windows.nwprinting.oob",
		registered: true,
	},
	{
		name:       "application/vnd.ms-windows.printerpairing",
		registered: true,
	},
	{
		name:       "application/vnd.ms-windows.wsd.oob",
		registered: true,
	},
	{
		name:       "application/vnd.ms-wmdrm.lic-chlg-req",
		registered: true,
	},
	{
		name:       "application/vnd.ms-wmdrm.lic-resp",
		registered: true,
	},
	{
		name:       "application/vnd.ms-wmdrm.meter-chlg-req",
		registered: true,
	},
	{
		name:       "application/vnd.ms-wmdrm.meter-resp",
		registered: true,
	},
	{
		name:       "application/vnd.ms-word.document.macroenabled.12",
		registered: false,
		extensions: []string{
			"docm",
		},
	},
	{
		name:       "application/vnd.ms-word.template.macroEnabled.12",
		registered: true,
		extensions: []string{
			"dotm",
		},
	},
	{
		name:       "application/vnd.ms-works",
		registered: true,
		extensions: []string{
			"wps", "wks", "wcm", "wdb",
		},
	},
	{
		name:       "application/vnd.ms-wpl",
		registered: true,
		extensions: []string{
			"wpl",
		},
	},
	{
		name:       "application/vnd.ms-xpsdocument",
		registered: true,
		extensions: []string{
			"xps",
		},
	},
	{
		name:       "application/vnd.msa-disk-image",
		registered: true,
	},
	{
		name:       "application/vnd.mseq",
		registered: true,
		extensions: []string{
			"mseq",
		},
	},
	{
		name:       "application/vnd.msign",
		registered: true,
	},
	{
		name:       "application/vnd.multiad.creator",
		registered: true,
	},
	{
		name:       "application/vnd.multiad.creator.cif",
		registered: true,
	},
	{
		name:       "application/vnd.music-niff",
		registered: true,
	},
	{
		name:       "application/vnd.musician",
		registered: true,
		extensions: []string{
			"mus",
		},
	},
	{
		name:       "application/vnd.muvee.style",
		registered: true,
		extensions: []string{
			"msty",
		},
	},
	{
		name:       "application/vnd.mynfc",
		registered: true,
		extensions: []string{
			"taglet",
		},
	},
	{
		name:       "application/vnd.nacamar.ybrid+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.ncd.control",
		registered: true,
	},
	{
		name:       "application/vnd.ncd.reference",
		registered: true,
	},
	{
		name:       "application/vnd.nearst.inv+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.nebumind.line",
		registered: true,
	},
	{
		name:       "application/vnd.nervana",
		registered: true,
	},
	{
		name:       "application/vnd.netfpx",
		registered: true,
	},
	{
		name:       "application/vnd.neurolanguage.nlu",
		registered: true,
		extensions: []string{
			"nlu",
		},
	},
	{
		name:       "application/vnd.nimn",
		registered: true,
	},
	{
		name:       "application/vnd.nintendo.nitro.rom",
		registered: true,
	},
	{
		name:       "application/vnd.nintendo.snes.rom",
		registered: true,
	},
	{
		name:       "application/vnd.nitf",
		registered: true,
		extensions: []string{
			"ntf", "nitf",
		},
	},
	{
		name:       "application/vnd.noblenet-directory",
		registered: true,
		extensions: []string{
			"nnd",
		},
	},
	{
		name:       "application/vnd.noblenet-sealer",
		registered: true,
		extensions: []string{
			"nns",
		},
	},
	{
		name:       "application/vnd.noblenet-web",
		registered: true,
		extensions: []string{
			"nnw",
		},
	},
	{
		name:       "application/vnd.nokia.catalogs",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.configuration-message",
		registered: false,
		extensions: []string{
			"ncm",
		},
	},
	{
		name:       "application/vnd.nokia.conml+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.conml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.iSDS-radio-presets",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.iptv.config+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.landmark+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.landmark+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.landmarkcollection+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.n-gage.ac+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.n-gage.data",
		registered: true,
		extensions: []string{
			"ngdat",
		},
	},
	{
		name:       "application/vnd.nokia.n-gage.symbian.install",
		registered: true,
		extensions: []string{
			"n-gage",
		},
	},
	{
		name:       "application/vnd.nokia.ncd",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.pcd+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.pcd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.nokia.radio-preset",
		registered: true,
		extensions: []string{
			"rpst",
		},
	},
	{
		name:       "application/vnd.nokia.radio-presets",
		registered: true,
		extensions: []string{
			"rpss",
		},
	},
	{
		name:       "application/vnd.nokia.ringing-tone",
		registered: false,
		extensions: []string{
			"rng",
		},
	},
	{
		name:       "application/vnd.novadigm.EDX",
		registered: true,
		extensions: []string{
			"edx",
		},
	},
	{
		name:       "application/vnd.novadigm.EXT",
		registered: true,
		extensions: []string{
			"ext",
		},
	},
	{
		name:       "application/vnd.novadigm.edm",
		registered: false,
		extensions: []string{
			"edm",
		},
	},
	{
		name:       "application/vnd.ntt-local.content-share",
		registered: true,
	},
	{
		name:       "application/vnd.ntt-local.file-transfer",
		registered: true,
	},
	{
		name:       "application/vnd.ntt-local.ogw_remote-access",
		registered: true,
	},
	{
		name:       "application/vnd.ntt-local.sip-ta_remote",
		registered: true,
	},
	{
		name:       "application/vnd.ntt-local.sip-ta_tcp_stream",
		registered: true,
	},
	{
		name:       "application/vnd.oasis.opendocument.base",
		registered: true,
	},
	{
		name:       "application/vnd.oasis.opendocument.chart",
		registered: true,
		extensions: []string{
			"odc",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.chart-template",
		registered: true,
		extensions: []string{
			"otc",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.database",
		registered: true,
		extensions: []string{
			"odb",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.formula",
		registered: true,
		extensions: []string{
			"odf",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.formula-template",
		registered: true,
		extensions: []string{
			"odft",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.graphics",
		registered: true,
		extensions: []string{
			"odg",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.graphics-template",
		registered: true,
		extensions: []string{
			"otg",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.image",
		registered: true,
		extensions: []string{
			"odi",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.image-template",
		registered: true,
		extensions: []string{
			"oti",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.presentation",
		registered: true,
		extensions: []string{
			"odp",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.presentation-template",
		registered: true,
		extensions: []string{
			"otp",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.spreadsheet",
		registered: true,
		extensions: []string{
			"ods",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.spreadsheet-template",
		registered: true,
		extensions: []string{
			"ots",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.text",
		registered: true,
		extensions: []string{
			"odt",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.text-master",
		registered: true,
		extensions: []string{
			"odm", "otm",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.text-template",
		registered: true,
		extensions: []string{
			"ott",
		},
	},
	{
		name:       "application/vnd.oasis.opendocument.text-web",
		registered: true,
		extensions: []string{
			"oth",
		},
	},
	{
		name:       "application/vnd.obn",
		registered: true,
	},
	{
		name:       "application/vnd.ocf+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/vnd.oci.image.manifest.v1+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.oftn.l10n+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.contentaccessdownload+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.contentaccessstreaming+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.cspg-hexbinary",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.dae.svg+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.dae.xhtml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.mippvcontrolmessage+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.pae.gem",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.spdiscovery+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.spdlist+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.ueprofile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oipf.userprofile+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.olpc-sugar",
		registered: true,
		extensions: []string{
			"xo",
		},
	},
	{
		name:       "application/vnd.oma-scws-config",
		registered: true,
	},
	{
		name:       "application/vnd.oma-scws-http-request",
		registered: true,
	},
	{
		name:       "application/vnd.oma-scws-http-response",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.associated-procedure-parameter+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.drm-trigger+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.imd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.ltkm",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.notification+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.provisioningtrigger",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.sgboot",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.sgdd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.sgdu",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.simple-symbol-container",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.smartcard-trigger+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.sprov+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.bcast.stkm",
		registered: true,
	},
	{
		name:       "application/vnd.oma.cab-address-book+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.cab-feature-handler+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.cab-pcc+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.cab-subs-invite+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.cab-user-prefs+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.dcd",
		registered: true,
	},
	{
		name:       "application/vnd.oma.dcdc",
		registered: true,
	},
	{
		name:       "application/vnd.oma.dd2+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"dd2",
		},
	},
	{
		name:       "application/vnd.oma.drm.risd+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.group-usage-list+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.lwm2m+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/vnd.oma.lwm2m+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.oma.lwm2m+tlv",
		registered: true,
	},
	{
		name:       "application/vnd.oma.pal+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.poc.detailed-progress-report+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.poc.final-report+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.poc.groups+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.poc.invocation-descriptor+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.poc.optimized-progress-report+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.push",
		registered: true,
	},
	{
		name:       "application/vnd.oma.scidm.messages+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oma.xcap-directory+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.omads-email+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.omads-file+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.omads-folder+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.omaloc-supl-init",
		registered: true,
	},
	{
		name:       "application/vnd.onepager",
		registered: true,
	},
	{
		name:       "application/vnd.onepagertamp",
		registered: true,
	},
	{
		name:       "application/vnd.onepagertamx",
		registered: true,
	},
	{
		name:       "application/vnd.onepagertat",
		registered: true,
	},
	{
		name:       "application/vnd.onepagertatp",
		registered: true,
	},
	{
		name:       "application/vnd.onepagertatx",
		registered: true,
	},
	{
		name:       "application/vnd.onvif.metadata",
		registered: true,
	},
	{
		name:       "application/vnd.openblox.game+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openblox.game-binary",
		registered: true,
	},
	{
		name:       "application/vnd.openeye.oeb",
		registered: true,
	},
	{
		name:       "application/vnd.openofficeorg.extension",
		registered: false,
		extensions: []string{
			"oxt",
		},
	},
	{
		name:       "application/vnd.openstreetmap.data+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.opentimestamps.ots",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.custom-properties+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.customXmlProperties+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawing+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawingml.chart+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawingml.chartshapes+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawingml.diagramcolors+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawingml.diagramdata+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.drawingml.diagramstyle+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.extended-properties+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.commentauthors+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.comments+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.handoutmaster+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		registered: true,
		extensions: []string{
			"pptx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.presprops+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slide",
		registered: true,
		extensions: []string{
			"sldx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slide+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slideUpdateInfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slidelayout+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slideshow",
		registered: true,
		extensions: []string{
			"ppsx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.tags+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.template",
		registered: true,
		extensions: []string{
			"potx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.template.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.calcChain+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.connections+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.dialogsheet+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.externalLink+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheRecords+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotTable+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotcachedefinition+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.querytable+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.revisionheaders+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.revisionlog+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedstrings+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		registered: true,
		extensions: []string{
			"xlsx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.sheetMetadata+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.tableSingleCells+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.template",
		registered: true,
		extensions: []string{
			"xltx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.usernames+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.volatiledependencies+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.theme+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.themeoverride+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.vmlDrawing",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		registered: true,
		extensions: []string{
			"docx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.document.glossary+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.template",
		registered: true,
		extensions: []string{
			"dotx",
		},
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-officedocument.wordprocessingml.websettings+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/vnd.openxmlformats-package.core-properties+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-package.digital-signature-xmlsignature+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.openxmlformats-package.relationships+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oracle.resource+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.orange.indata",
		registered: true,
	},
	{
		name:       "application/vnd.osa.netdeploy",
		registered: true,
	},
	{
		name:       "application/vnd.osgeo.mapguide.package",
		registered: true,
		extensions: []string{
			"mgp",
		},
	},
	{
		name:       "application/vnd.osgi.bundle",
		registered: true,
	},
	{
		name:       "application/vnd.osgi.dp",
		registered: true,
		extensions: []string{
			"dp",
		},
	},
	{
		name:       "application/vnd.osgi.subsystem",
		registered: true,
		extensions: []string{
			"esa",
		},
	},
	{
		name:       "application/vnd.otps.ct-kip+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.oxli.countgraph",
		registered: true,
	},
	{
		name:       "application/vnd.pagerduty+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.palm",
		registered: true,
		extensions: []string{
			"pdb", "pqa", "oprc",
		},
	},
	{
		name:       "application/vnd.panoply",
		registered: true,
	},
	{
		name:       "application/vnd.paos.xml",
		registered: true,
	},
	{
		name:       "application/vnd.patentdive",
		registered: true,
	},
	{
		name:       "application/vnd.patientecommsdoc",
		registered: true,
	},
	{
		name:       "application/vnd.pawaafile",
		registered: true,
		extensions: []string{
			"paw",
		},
	},
	{
		name:       "application/vnd.pcos",
		registered: true,
	},
	{
		name:       "application/vnd.pg.format",
		registered: true,
		extensions: []string{
			"str",
		},
	},
	{
		name:       "application/vnd.pg.osasli",
		registered: true,
		extensions: []string{
			"ei6",
		},
	},
	{
		name:       "application/vnd.piaccess.application-licence",
		registered: true,
	},
	{
		name:       "application/vnd.picsel",
		registered: true,
		extensions: []string{
			"efif",
		},
	},
	{
		name:       "application/vnd.pmi.widget",
		registered: true,
		extensions: []string{
			"wg",
		},
	},
	{
		name:       "application/vnd.poc.group-advertisement+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.pocketlearn",
		registered: true,
		extensions: []string{
			"plf",
		},
	},
	{
		name:       "application/vnd.powerbuilder6",
		registered: true,
		extensions: []string{
			"pbd",
		},
	},
	{
		name:       "application/vnd.powerbuilder6-s",
		registered: true,
	},
	{
		name:       "application/vnd.powerbuilder7",
		registered: true,
	},
	{
		name:       "application/vnd.powerbuilder7-s",
		registered: true,
	},
	{
		name:       "application/vnd.powerbuilder75",
		registered: true,
	},
	{
		name:       "application/vnd.powerbuilder75-s",
		registered: true,
	},
	{
		name:       "application/vnd.preminet",
		registered: true,
	},
	{
		name:       "application/vnd.previewsystems.box",
		registered: true,
		extensions: []string{
			"box",
		},
	},
	{
		name:       "application/vnd.proteus.magazine",
		registered: true,
		extensions: []string{
			"mgz",
		},
	},
	{
		name:       "application/vnd.psfs",
		registered: true,
	},
	{
		name:       "application/vnd.publishare-delta-tree",
		registered: true,
		extensions: []string{
			"qps",
		},
	},
	{
		name:       "application/vnd.pvi.ptid1",
		registered: true,
		extensions: []string{
			"ptid",
		},
	},
	{
		name:       "application/vnd.pwg-multiplexed",
		registered: true,
	},
	{
		name:       "application/vnd.pwg-xhtml-print+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.qualcomm.brew-app-res",
		registered: true,
	},
	{
		name:       "application/vnd.quarantainenet",
		registered: true,
	},
	{
		name:       "application/vnd.quobject-quoxdocument",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.moml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-audit+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-audit-conf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-audit-conn+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-audit-dialog+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-audit-stream+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-conf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog-base+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog-fax-detect+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog-fax-sendrecv+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog-group+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog-speech+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.radisys.msml-dialog-transform+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.rainstor.data",
		registered: true,
	},
	{
		name:       "application/vnd.rapid",
		registered: true,
	},
	{
		name:       "application/vnd.rar",
		registered: true,
	},
	{
		name:       "application/vnd.realvnc.bed",
		registered: true,
		extensions: []string{
			"bed",
		},
	},
	{
		name:       "application/vnd.recordare.musicxml",
		registered: true,
		extensions: []string{
			"mxl",
		},
	},
	{
		name:       "application/vnd.recordare.musicxml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"musicxml",
		},
	},
	{
		name:       "application/vnd.renlearn.rlprint",
		registered: false,
	},
	{
		name:       "application/vnd.resilient.logic",
		registered: true,
	},
	{
		name:       "application/vnd.restful+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.rig.cryptonote",
		registered: true,
		extensions: []string{
			"cryptonote",
		},
	},
	{
		name:       "application/vnd.rim.cod",
		registered: false,
		extensions: []string{
			"cod",
		},
	},
	{
		name:       "application/vnd.rn-realmedia",
		registered: false,
		extensions: []string{
			"rm",
		},
	},
	{
		name:       "application/vnd.rn-realmedia-vbr",
		registered: false,
		extensions: []string{
			"rmvb",
		},
	},
	{
		name:       "application/vnd.rn-realplayer",
		registered: false,
		extensions: []string{
			"rnx",
		},
	},
	{
		name:       "application/vnd.route66.link66+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"link66",
		},
	},
	{
		name:       "application/vnd.rs-274x",
		registered: true,
	},
	{
		name:       "application/vnd.ruckus.download",
		registered: true,
	},
	{
		name:       "application/vnd.s3sms",
		registered: true,
	},
	{
		name:       "application/vnd.sailingtracker.track",
		registered: true,
		extensions: []string{
			"st",
		},
	},
	{
		name:       "application/vnd.sar",
		registered: true,
	},
	{
		name:       "application/vnd.sbm.cid",
		registered: true,
	},
	{
		name:       "application/vnd.sbm.mid2",
		registered: true,
	},
	{
		name:       "application/vnd.scribus",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.3df",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.csf",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.doc",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.eml",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.mht",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.net",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.ppt",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.tiff",
		registered: true,
	},
	{
		name:       "application/vnd.sealed.xls",
		registered: true,
	},
	{
		name:       "application/vnd.sealedmedia.softseal.html",
		registered: true,
	},
	{
		name:       "application/vnd.sealedmedia.softseal.pdf",
		registered: true,
	},
	{
		name:       "application/vnd.seemail",
		registered: true,
		extensions: []string{
			"see",
		},
	},
	{
		name:       "application/vnd.seis+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.sema",
		registered: true,
		extensions: []string{
			"sema",
		},
	},
	{
		name:       "application/vnd.semd",
		registered: true,
		extensions: []string{
			"semd",
		},
	},
	{
		name:       "application/vnd.semf",
		registered: true,
		extensions: []string{
			"semf",
		},
	},
	{
		name:       "application/vnd.shade-save-file",
		registered: true,
	},
	{
		name:       "application/vnd.shana.informed.formdata",
		registered: true,
		extensions: []string{
			"ifm",
		},
	},
	{
		name:       "application/vnd.shana.informed.formtemplate",
		registered: true,
		extensions: []string{
			"itp",
		},
	},
	{
		name:       "application/vnd.shana.informed.interchange",
		registered: true,
		extensions: []string{
			"iif",
		},
	},
	{
		name:       "application/vnd.shana.informed.package",
		registered: true,
		extensions: []string{
			"ipk",
		},
	},
	{
		name:       "application/vnd.shootproof+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.shopkick+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.shp",
		registered: true,
	},
	{
		name:       "application/vnd.shx",
		registered: true,
	},
	{
		name:       "application/vnd.sigrok.session",
		registered: true,
	},
	{
		name:       "application/vnd.siren+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.smaf",
		registered: true,
		extensions: []string{
			"mmf",
		},
	},
	{
		name:       "application/vnd.smart.notebook",
		registered: true,
	},
	{
		name:       "application/vnd.smart.teacher",
		registered: true,
		extensions: []string{
			"teacher",
		},
	},
	{
		name:       "application/vnd.snesdev-page-table",
		registered: true,
	},
	{
		name:       "application/vnd.software602.filler.form+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.software602.filler.form-xml-zip",
		registered: true,
	},
	{
		name:       "application/vnd.solent.sdkm+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"sdkm", "sdkd",
		},
	},
	{
		name:       "application/vnd.spotfire.dxp",
		registered: true,
		extensions: []string{
			"dxp",
		},
	},
	{
		name:       "application/vnd.spotfire.sfs",
		registered: true,
		extensions: []string{
			"sfs",
		},
	},
	{
		name:       "application/vnd.sqlite3",
		registered: true,
	},
	{
		name:       "application/vnd.sss-cod",
		registered: true,
	},
	{
		name:       "application/vnd.sss-dtf",
		registered: true,
	},
	{
		name:       "application/vnd.sss-ntf",
		registered: true,
	},
	{
		name:       "application/vnd.stardivision.calc",
		registered: false,
		extensions: []string{
			"sdc",
		},
	},
	{
		name:       "application/vnd.stardivision.draw",
		registered: false,
		extensions: []string{
			"sda",
		},
	},
	{
		name:       "application/vnd.stardivision.impress",
		registered: false,
		extensions: []string{
			"sdd", "sdp",
		},
	},
	{
		name:       "application/vnd.stardivision.math",
		registered: false,
		extensions: []string{
			"smf",
		},
	},
	{
		name:       "application/vnd.stardivision.writer",
		registered: false,
		extensions: []string{
			"sdw", "vor",
		},
	},
	{
		name:       "application/vnd.stardivision.writer-global",
		registered: false,
		extensions: []string{
			"sgl",
		},
	},
	{
		name:       "application/vnd.stepmania.package",
		registered: true,
		extensions: []string{
			"smzip",
		},
	},
	{
		name:       "application/vnd.stepmania.stepchart",
		registered: true,
		extensions: []string{
			"sm",
		},
	},
	{
		name:       "application/vnd.street-stream",
		registered: true,
	},
	{
		name:       "application/vnd.sun.wadl+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.sun.xml.calc",
		registered: false,
		extensions: []string{
			"sxc",
		},
	},
	{
		name:       "application/vnd.sun.xml.calc.template",
		registered: false,
		extensions: []string{
			"stc",
		},
	},
	{
		name:       "application/vnd.sun.xml.draw",
		registered: false,
		extensions: []string{
			"sxd",
		},
	},
	{
		name:       "application/vnd.sun.xml.draw.template",
		registered: false,
		extensions: []string{
			"std",
		},
	},
	{
		name:       "application/vnd.sun.xml.impress",
		registered: false,
		extensions: []string{
			"sxi",
		},
	},
	{
		name:       "application/vnd.sun.xml.impress.template",
		registered: false,
		extensions: []string{
			"sti",
		},
	},
	{
		name:       "application/vnd.sun.xml.math",
		registered: false,
		extensions: []string{
			"sxm",
		},
	},
	{
		name:       "application/vnd.sun.xml.writer",
		registered: false,
		extensions: []string{
			"sxw",
		},
	},
	{
		name:       "application/vnd.sun.xml.writer.global",
		registered: false,
		extensions: []string{
			"sxg",
		},
	},
	{
		name:       "application/vnd.sun.xml.writer.template",
		registered: false,
		extensions: []string{
			"stw",
		},
	},
	{
		name:       "application/vnd.sus-calendar",
		registered: true,
		extensions: []string{
			"sus", "susp",
		},
	},
	{
		name:       "application/vnd.svd",
		registered: true,
		extensions: []string{
			"svd",
		},
	},
	{
		name:       "application/vnd.swiftview-ics",
		registered: true,
	},
	{
		name:       "application/vnd.sybyl.mol2",
		registered: true,
	},
	{
		name:       "application/vnd.sycle+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.syft+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.symbian.install",
		registered: false,
		extensions: []string{
			"sis", "sisx",
		},
	},
	{
		name:       "application/vnd.syncml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xsm",
		},
	},
	{
		name:       "application/vnd.syncml.dm+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
		extensions: []string{
			"bdm",
		},
	},
	{
		name:       "application/vnd.syncml.dm+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xdm",
		},
	},
	{
		name:       "application/vnd.syncml.dm.notification",
		registered: true,
	},
	{
		name:       "application/vnd.syncml.dmddf+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.syncml.dmddf+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.syncml.dmtnds+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.syncml.dmtnds+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.syncml.ds.notification",
		registered: true,
	},
	{
		name:       "application/vnd.tableschema+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.tao.intent-module-archive",
		registered: true,
		extensions: []string{
			"tao",
		},
	},
	{
		name:       "application/vnd.tcpdump.pcap",
		registered: true,
		extensions: []string{
			"pcap", "cap", "dmp",
		},
	},
	{
		name:       "application/vnd.think-cell.ppttc+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.tmd.mediaflex.api+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.tml",
		registered: true,
	},
	{
		name:       "application/vnd.tmobile-livetv",
		registered: true,
		extensions: []string{
			"tmo",
		},
	},
	{
		name:       "application/vnd.tri.onesource",
		registered: true,
	},
	{
		name:       "application/vnd.trid.tpt",
		registered: true,
		extensions: []string{
			"tpt",
		},
	},
	{
		name:       "application/vnd.triscape.mxs",
		registered: true,
		extensions: []string{
			"mxs",
		},
	},
	{
		name:       "application/vnd.trueapp",
		registered: true,
		extensions: []string{
			"tra",
		},
	},
	{
		name:       "application/vnd.truedoc",
		registered: true,
	},
	{
		name:       "application/vnd.tve-trigger",
		registered: false,
	},
	{
		name:       "application/vnd.ubisoft.webplayer",
		registered: true,
	},
	{
		name:       "application/vnd.ufdl",
		registered: true,
		extensions: []string{
			"ufd", "ufdl",
		},
	},
	{
		name:       "application/vnd.uiq.theme",
		registered: true,
		extensions: []string{
			"utz",
		},
	},
	{
		name:       "application/vnd.umajin",
		registered: true,
		extensions: []string{
			"umj",
		},
	},
	{
		name:       "application/vnd.unity",
		registered: true,
		extensions: []string{
			"unityweb",
		},
	},
	{
		name:       "application/vnd.uoml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"uoml",
		},
	},
	{
		name:       "application/vnd.uplanet.alert",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.alert-wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.bearer-choice",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.bearer-choice-wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.cacheop",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.cacheop-wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.channel",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.channel-wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.list",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.list-wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.listcmd",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.listcmd-wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.uplanet.signal",
		registered: true,
	},
	{
		name:       "application/vnd.uri-map",
		registered: true,
	},
	{
		name:       "application/vnd.valve.source.material",
		registered: true,
	},
	{
		name:       "application/vnd.vcx",
		registered: true,
		extensions: []string{
			"vcx",
		},
	},
	{
		name:       "application/vnd.vd-study",
		registered: true,
	},
	{
		name:       "application/vnd.vectorworks",
		registered: true,
	},
	{
		name:       "application/vnd.vel+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.verimatrix.vcas",
		registered: true,
	},
	{
		name:       "application/vnd.veritone.aion+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.veryant.thin",
		registered: true,
	},
	{
		name:       "application/vnd.ves.encrypted",
		registered: true,
	},
	{
		name:       "application/vnd.vidsoft.vidconference",
		registered: true,
	},
	{
		name:       "application/vnd.visio",
		registered: true,
		extensions: []string{
			"vsd", "vst", "vss", "vsw",
		},
	},
	{
		name:       "application/vnd.visionary",
		registered: true,
		extensions: []string{
			"vis",
		},
	},
	{
		name:       "application/vnd.vividence.scriptfile",
		registered: true,
	},
	{
		name:       "application/vnd.vsf",
		registered: true,
		extensions: []string{
			"vsf",
		},
	},
	{
		name:       "application/vnd.wap.sic",
		registered: true,
		extensions: []string{
			"sic",
		},
	},
	{
		name:       "application/vnd.wap.slc",
		registered: true,
		extensions: []string{
			"slc",
		},
	},
	{
		name:       "application/vnd.wap.wbxml",
		registered: true,
		extensions: []string{
			"wbxml",
		},
	},
	{
		name:       "application/vnd.wap.wmlc",
		registered: true,
		extensions: []string{
			"wmlc",
		},
	},
	{
		name:       "application/vnd.wap.wmlscriptc",
		registered: true,
		extensions: []string{
			"wmlsc",
		},
	},
	{
		name:       "application/vnd.wasmflow.wafl",
		registered: true,
	},
	{
		name:       "application/vnd.webturbo",
		registered: true,
		extensions: []string{
			"wtb",
		},
	},
	{
		name:       "application/vnd.wfa.dpp",
		registered: true,
	},
	{
		name:       "application/vnd.wfa.p2p",
		registered: true,
	},
	{
		name:       "application/vnd.wfa.wsc",
		registered: true,
	},
	{
		name:       "application/vnd.windows.devicepairing",
		registered: true,
	},
	{
		name:       "application/vnd.wmc",
		registered: true,
	},
	{
		name:       "application/vnd.wmf.bootstrap",
		registered: true,
	},
	{
		name:       "application/vnd.wolfram.mathematica",
		registered: true,
	},
	{
		name:       "application/vnd.wolfram.mathematica.package",
		registered: true,
	},
	{
		name:       "application/vnd.wolfram.player",
		registered: true,
		extensions: []string{
			"nbp",
		},
	},
	{
		name:       "application/vnd.wordperfect",
		registered: true,
		extensions: []string{
			"wpd",
		},
	},
	{
		name:       "application/vnd.wqd",
		registered: true,
		extensions: []string{
			"wqd",
		},
	},
	{
		name:       "application/vnd.wrq-hp3000-labelled",
		registered: true,
	},
	{
		name:       "application/vnd.wt.stf",
		registered: true,
		extensions: []string{
			"stf",
		},
	},
	{
		name:       "application/vnd.wv.csp+wbxml",
		format:     "application/vnd.wap.wbxml",
		registered: true,
	},
	{
		name:       "application/vnd.wv.csp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.wv.ssp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.xacml+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vnd.xara",
		registered: true,
		extensions: []string{
			"xar", "web",
		},
	},
	{
		name:       "application/vnd.xfdl",
		registered: true,
		extensions: []string{
			"xfdl",
		},
	},
	{
		name:       "application/vnd.xfdl.webform",
		registered: true,
	},
	{
		name:       "application/vnd.xmi+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/vnd.xmpie.cpkg",
		registered: true,
	},
	{
		name:       "application/vnd.xmpie.dpkg",
		registered: true,
	},
	{
		name:       "application/vnd.xmpie.plan",
		registered: true,
	},
	{
		name:       "application/vnd.xmpie.ppkg",
		registered: true,
	},
	{
		name:       "application/vnd.xmpie.xlim",
		registered: true,
	},
	{
		name:       "application/vnd.yamaha.hv-dic",
		registered: true,
		extensions: []string{
			"hvd",
		},
	},
	{
		name:       "application/vnd.yamaha.hv-script",
		registered: true,
		extensions: []string{
			"hvs",
		},
	},
	{
		name:       "application/vnd.yamaha.hv-voice",
		registered: true,
		extensions: []string{
			"hvp",
		},
	},
	{
		name:       "application/vnd.yamaha.openscoreformat",
		registered: true,
		extensions: []string{
			"osf",
		},
	},
	{
		name:       "application/vnd.yamaha.openscoreformat.osfpvg+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"osfpvg",
		},
	},
	{
		name:       "application/vnd.yamaha.remote-setup",
		registered: true,
	},
	{
		name:       "application/vnd.yamaha.smaf-audio",
		registered: true,
		extensions: []string{
			"saf",
		},
	},
	{
		name:       "application/vnd.yamaha.smaf-phrase",
		registered: true,
		extensions: []string{
			"spf",
		},
	},
	{
		name:       "application/vnd.yamaha.through-ngn",
		registered: true,
	},
	{
		name:       "application/vnd.yamaha.tunnel-udpencap",
		registered: true,
	},
	{
		name:       "application/vnd.yaoweme",
		registered: true,
	},
	{
		name:       "application/vnd.yellowriver-custom-menu",
		registered: true,
		extensions: []string{
			"cmp",
		},
	},
	{
		name:       "application/vnd.youtube.yt",
		registered: true,
	},
	{
		name:       "application/vnd.zul",
		registered: true,
		extensions: []string{
			"zir", "zirz",
		},
	},
	{
		name:       "application/vnd.zzazz.deck+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"zaz",
		},
	},
	{
		name:       "application/vocaltec-media-desc",
		registered: false,
		extensions: []string{
			"vmd",
		},
	},
	{
		name:       "application/vocaltec-media-file",
		registered: false,
		extensions: []string{
			"vmf",
		},
	},
	{
		name:       "application/voicexml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"vxml",
		},
	},
	{
		name:       "application/voucher-cms+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/vq-rtcpxr",
		registered: true,
	},
	{
		name:       "application/wasm",
		registered: true,
	},
	{
		name:       "application/watcherinfo+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/webpush-options+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/whoispp-query",
		registered: true,
	},
	{
		name:       "application/whoispp-response",
		registered: true,
	},
	{
		name:       "application/widget",
		registered: true,
		extensions: []string{
			"wgt",
		},
	},
	{
		name:       "application/winhlp",
		registered: false,
		extensions: []string{
			"hlp",
		},
	},
	{
		name:       "application/wita",
		registered: true,
	},
	{
		name:       "application/wordperfect",
		registered: false,
		extensions: []string{
			"wp", "wp5", "wp6", "wpd",
		},
	},
	{
		name:       "application/wordperfect5.1",
		registered: true,
		extensions: []string{
			"wp5",
		},
	},
	{
		name:       "application/wordperfect6.0",
		registered: false,
		extensions: []string{
			"w60", "wp5",
		},
	},
	{
		name:       "application/wordperfect6.1",
		registered: false,
		extensions: []string{
			"w61",
		},
	},
	{
		name:       "application/wsdl+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"wsdl",
		},
	},
	{
		name:       "application/wspolicy+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"wspolicy",
		},
	},
	{
		name:       "application/x-123",
		registered: false,
		extensions: []string{
			"wk1", "wk",
		},
	},
	{
		name:       "application/x-7z-compressed",
		registered: false,
		extensions: []string{
			"7z",
		},
	},
	{
		name:       "application/x-abiword",
		registered: false,
		extensions: []string{
			"abw",
		},
	},
	{
		name:       "application/x-ace-compressed",
		registered: false,
		extensions: []string{
			"ace",
		},
	},
	{
		name:       "application/x-aim",
		registered: false,
		extensions: []string{
			"aim",
		},
	},
	{
		name:       "application/x-amf",
		registered: false,
	},
	{
		name:       "application/x-apple-diskimage",
		registered: false,
		extensions: []string{
			"dmg",
		},
	},
	{
		name:       "application/x-authorware-bin",
		registered: false,
		extensions: []string{
			"aab", "x32", "u32", "vox",
		},
	},
	{
		name:       "application/x-authorware-map",
		registered: false,
		extensions: []string{
			"aam",
		},
	},
	{
		name:       "application/x-authorware-seg",
		registered: false,
		extensions: []string{
			"aas",
		},
	},
	{
		name:       "application/x-bcpio",
		registered: false,
		extensions: []string{
			"bcpio",
		},
	},
	{
		name:       "application/x-binary",
		registered: false,
		extensions: []string{
			"bin",
		},
	},
	{
		name:       "application/x-binhex40",
		registered: false,
		extensions: []string{
			"hqx",
		},
	},
	{
		name:       "application/x-bittorrent",
		registered: false,
		extensions: []string{
			"torrent",
		},
	},
	{
		name:       "application/x-blorb",
		registered: false,
		extensions: []string{
			"blb", "blorb",
		},
	},
	{
		name:       "application/x-bsh",
		registered: false,
		extensions: []string{
			"bsh", "sh", "shar",
		},
	},
	{
		name:       "application/x-bytecode.elisp",
		registered: false,
		extensions: []string{
			"elc",
		},
	},
	{
		name:       "application/x-bytecode.python",
		registered: false,
		extensions: []string{
			"pyc",
		},
	},
	{
		name:       "application/x-bzip",
		registered: false,
		extensions: []string{
			"bz",
		},
	},
	{
		name:       "application/x-bzip2",
		registered: false,
		extensions: []string{
			"bz2", "boz",
		},
	},
	{
		name:       "application/x-cbr",
		registered: false,
		extensions: []string{
			"cbr", "cba", "cbt", "cbz", "cb7",
		},
	},
	{
		name:       "application/x-cdf",
		registered: false,
		extensions: []string{
			"cdf",
		},
	},
	{
		name:       "application/x-cdlink",
		registered: false,
		extensions: []string{
			"vcd",
		},
	},
	{
		name:       "application/x-cfs-compressed",
		registered: false,
		extensions: []string{
			"cfs",
		},
	},
	{
		name:       "application/x-chat",
		registered: false,
		extensions: []string{
			"chat", "cha",
		},
	},
	{
		name:       "application/x-chess-pgn",
		registered: false,
		extensions: []string{
			"pgn",
		},
	},
	{
		name:       "application/x-chm",
		registered: false,
		extensions: []string{
			"chm",
		},
	},
	{
		name:       "application/x-chrome-extension",
		registered: false,
		extensions: []string{
			"crx",
		},
	},
	{
		name:       "application/x-cmu-raster",
		registered: false,
		extensions: []string{
			"ras",
		},
	},
	{
		name:       "application/x-cocoa",
		registered: false,
		extensions: []string{
			"cco",
		},
	},
	{
		name:       "application/x-compactpro",
		registered: false,
		extensions: []string{
			"cpt",
		},
	},
	{
		name:       "application/x-compress",
		registered: false,
		extensions: []string{
			"z",
		},
	},
	{
		name:       "application/x-compressed",
		registered: false,
		extensions: []string{
			"gz", "tgz", "z", "zip",
		},
	},
	{
		name:       "application/x-conference",
		registered: false,
		extensions: []string{
			"nsc",
		},
	},
	{
		name:       "application/x-core",
		registered: false,
	},
	{
		name:       "application/x-cpio",
		registered: false,
		extensions: []string{
			"cpio",
		},
	},
	{
		name:       "application/x-cpt",
		registered: false,
		extensions: []string{
			"cpt",
		},
	},
	{
		name:       "application/x-csh",
		registered: false,
		extensions: []string{
			"csh",
		},
	},
	{
		name:       "application/x-debian-package",
		registered: false,
		extensions: []string{
			"deb", "udeb",
		},
	},
	{
		name:       "application/x-deepv",
		registered: false,
		extensions: []string{
			"deepv",
		},
	},
	{
		name:       "application/x-dgc-compressed",
		registered: false,
		extensions: []string{
			"dgc",
		},
	},
	{
		name:       "application/x-director",
		registered: false,
		extensions: []string{
			"dir", "dcr", "dxr", "cst", "cct", "cxt", "w3d", "fgd", "swa",
		},
	},
	{
		name:       "application/x-dms",
		registered: false,
		extensions: []string{
			"dms",
		},
	},
	{
		name:       "application/x-doom",
		registered: false,
		extensions: []string{
			"wad",
		},
	},
	{
		name:       "application/x-dtbncx+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"ncx",
		},
	},
	{
		name:       "application/x-dtbook+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"dtb",
		},
	},
	{
		name:       "application/x-dtbresource+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"res",
		},
	},
	{
		name:       "application/x-dvi",
		registered: false,
		extensions: []string{
			"dvi",
		},
	},
	{
		name:       "application/x-elc",
		registered: false,
		extensions: []string{
			"elc",
		},
	},
	{
		name:       "application/x-envoy",
		registered: false,
		extensions: []string{
			"env", "evy",
		},
	},
	{
		name:       "application/x-esrehber",
		registered: false,
		extensions: []string{
			"es",
		},
	},
	{
		name:       "application/x-eva",
		registered: false,
		extensions: []string{
			"eva",
		},
	},
	{
		name:       "application/x-excel",
		registered: false,
		extensions: []string{
			"xla", "xlb", "xlc", "xld", "xlk", "xll", "xlm", "xls", "xlt", "xlv", "xlw",
		},
	},
	{
		name:       "application/x-executable",
		registered: false,
	},
	{
		name:       "application/x-flac",
		registered: false,
		extensions: []string{
			"flac",
		},
	},
	{
		name:       "application/x-font",
		registered: false,
		extensions: []string{
			"pfa", "pfb", "gsf", "pcf", "pcf.z",
		},
	},
	{
		name:       "application/x-font-bdf",
		registered: false,
		extensions: []string{
			"bdf",
		},
	},
	{
		name:       "application/x-font-dos",
		registered: false,
	},
	{
		name:       "application/x-font-framemaker",
		registered: false,
	},
	{
		name:       "application/x-font-ghostscript",
		registered: false,
		extensions: []string{
			"gsf",
		},
	},
	{
		name:       "application/x-font-libgrx",
		registered: false,
	},
	{
		name:       "application/x-font-linux-psf",
		registered: false,
		extensions: []string{
			"psf",
		},
	},
	{
		name:       "application/x-font-otf",
		registered: false,
		extensions: []string{
			"otf",
		},
	},
	{
		name:       "application/x-font-pcf",
		registered: false,
		extensions: []string{
			"pcf",
		},
	},
	{
		name:       "application/x-font-snf",
		registered: false,
		extensions: []string{
			"snf",
		},
	},
	{
		name:       "application/x-font-speedo",
		registered: false,
	},
	{
		name:       "application/x-font-sunos-news",
		registered: false,
	},
	{
		name:       "application/x-font-ttf",
		registered: false,
		extensions: []string{
			"ttf", "ttc",
		},
	},
	{
		name:       "application/x-font-type1",
		registered: false,
		extensions: []string{
			"pfa", "pfb", "pfm", "afm",
		},
	},
	{
		name:       "application/x-font-vfont",
		registered: false,
	},
	{
		name:       "application/x-font-woff",
		registered: false,
		extensions: []string{
			"woff",
		},
	},
	{
		name:       "application/x-frame",
		registered: false,
		extensions: []string{
			"mif",
		},
	},
	{
		name:       "application/x-freearc",
		registered: false,
		extensions: []string{
			"arc",
		},
	},
	{
		name:       "application/x-freelance",
		registered: false,
		extensions: []string{
			"pre",
		},
	},
	{
		name:       "application/x-futuresplash",
		registered: false,
		extensions: []string{
			"spl",
		},
	},
	{
		name:       "application/x-gca-compressed",
		registered: false,
		extensions: []string{
			"gca",
		},
	},
	{
		name:       "application/x-glulx",
		registered: false,
		extensions: []string{
			"ulx",
		},
	},
	{
		name:       "application/x-gnumeric",
		registered: false,
		extensions: []string{
			"gnumeric",
		},
	},
	{
		name:       "application/x-go-sgf",
		registered: false,
		extensions: []string{
			"sgf",
		},
	},
	{
		name:       "application/x-gramps-xml",
		registered: false,
		extensions: []string{
			"gramps",
		},
	},
	{
		name:       "application/x-graphing-calculator",
		registered: false,
		extensions: []string{
			"gcf",
		},
	},
	{
		name:       "application/x-gsp",
		registered: false,
		extensions: []string{
			"gsp",
		},
	},
	{
		name:       "application/x-gss",
		registered: false,
		extensions: []string{
			"gss",
		},
	},
	{
		name:       "application/x-gtar",
		registered: false,
		extensions: []string{
			"gtar", "tgz", "taz",
		},
	},
	{
		name:       "application/x-gzip",
		registered: false,
		extensions: []string{
			"gz", "gzip", "tgz",
		},
	},
	{
		name:       "application/x-hdf",
		registered: false,
		extensions: []string{
			"hdf",
		},
	},
	{
		name:       "application/x-helpfile",
		registered: false,
		extensions: []string{
			"help", "hlp",
		},
	},
	{
		name:       "application/x-httpd-imap",
		registered: false,
		extensions: []string{
			"imap",
		},
	},
	{
		name:       "application/x-httpd-php",
		registered: false,
		extensions: []string{
			"phtml", "pht", "php",
		},
	},
	{
		name:       "application/x-httpd-php-source",
		registered: false,
		extensions: []string{
			"phps",
		},
	},
	{
		name:       "application/x-httpd-php3",
		registered: false,
		extensions: []string{
			"php3",
		},
	},
	{
		name:       "application/x-httpd-php3-preprocessed",
		registered: false,
		extensions: []string{
			"php3p",
		},
	},
	{
		name:       "application/x-httpd-php4",
		registered: false,
		extensions: []string{
			"php4",
		},
	},
	{
		name:       "application/x-ica",
		registered: false,
		extensions: []string{
			"ica",
		},
	},
	{
		name:       "application/x-ima",
		registered: false,
		extensions: []string{
			"ima",
		},
	},
	{
		name:       "application/x-install-instructions",
		registered: false,
		extensions: []string{
			"install",
		},
	},
	{
		name:       "application/x-internet-signup",
		registered: false,
		extensions: []string{
			"ins", "isp",
		},
	},
	{
		name:       "application/x-internett-signup",
		registered: false,
		extensions: []string{
			"ins",
		},
	},
	{
		name:       "application/x-inventor",
		registered: false,
		extensions: []string{
			"iv",
		},
	},
	{
		name:       "application/x-ip2",
		registered: false,
		extensions: []string{
			"ip",
		},
	},
	{
		name:       "application/x-iphone",
		registered: false,
		extensions: []string{
			"iii",
		},
	},
	{
		name:       "application/x-iso9660-image",
		registered: false,
		extensions: []string{
			"iso",
		},
	},
	{
		name:       "application/x-java-applet",
		registered: false,
	},
	{
		name:       "application/x-java-archive",
		registered: false,
		extensions: []string{
			"jar",
		},
	},
	{
		name:       "application/x-java-bean",
		registered: false,
	},
	{
		name:       "application/x-java-class",
		registered: false,
		extensions: []string{
			"class",
		},
	},
	{
		name:       "application/x-java-commerce",
		registered: false,
		extensions: []string{
			"jcm",
		},
	},
	{
		name:       "application/x-java-jnlp-file",
		registered: false,
		extensions: []string{
			"jnlp",
		},
	},
	{
		name:       "application/x-java-serialized-object",
		registered: false,
		extensions: []string{
			"ser",
		},
	},
	{
		name:       "application/x-java-vm",
		registered: false,
		extensions: []string{
			"class",
		},
	},
	{
		name:       "application/x-javascript",
		registered: false,
		extensions: []string{
			"js",
		},
	},
	{
		name:       "application/x-kchart",
		registered: false,
		extensions: []string{
			"chrt",
		},
	},
	{
		name:       "application/x-kdelnk",
		registered: false,
	},
	{
		name:       "application/x-killustrator",
		registered: false,
		extensions: []string{
			"kil",
		},
	},
	{
		name:       "application/x-koan",
		registered: false,
		extensions: []string{
			"skd", "skm", "skp", "skt",
		},
	},
	{
		name:       "application/x-kpresenter",
		registered: false,
		extensions: []string{
			"kpr", "kpt",
		},
	},
	{
		name:       "application/x-ksh",
		registered: false,
		extensions: []string{
			"ksh",
		},
	},
	{
		name:       "application/x-kspread",
		registered: false,
		extensions: []string{
			"ksp",
		},
	},
	{
		name:       "application/x-kword",
		registered: false,
		extensions: []string{
			"kwd", "kwt",
		},
	},
	{
		name:       "application/x-latex",
		registered: false,
		extensions: []string{
			"latex", "ltx",
		},
	},
	{
		name:       "application/x-lha",
		registered: false,
		extensions: []string{
			"lha",
		},
	},
	{
		name:       "application/x-lisp",
		registered: false,
		extensions: []string{
			"lsp",
		},
	},
	{
		name:       "application/x-livescreen",
		registered: false,
		extensions: []string{
			"ivy",
		},
	},
	{
		name:       "application/x-lotus",
		registered: false,
		extensions: []string{
			"wq1",
		},
	},
	{
		name:       "application/x-lotusscreencam",
		registered: false,
		extensions: []string{
			"scm",
		},
	},
	{
		name:       "application/x-lua-bytecode",
		registered: false,
		extensions: []string{
			"luac",
		},
	},
	{
		name:       "application/x-lzh",
		registered: false,
		extensions: []string{
			"lzh",
		},
	},
	{
		name:       "application/x-lzh-compressed",
		registered: false,
		extensions: []string{
			"lzh", "lha",
		},
	},
	{
		name:       "application/x-lzx",
		registered: false,
		extensions: []string{
			"lzx",
		},
	},
	{
		name:       "application/x-mac-binhex40",
		registered: false,
		extensions: []string{
			"hqx",
		},
	},
	{
		name:       "application/x-macbinary",
		registered: false,
		extensions: []string{
			"bin",
		},
	},
	{
		name:       "application/x-magic-cap-package-1.0",
		registered: false,
		extensions: []string{
			"mc$",
		},
	},
	{
		name:       "application/x-maker",
		registered: false,
		extensions: []string{
			"frm", "maker", "frame", "fm", "fb", "book", "fbdoc",
		},
	},
	{
		name:       "application/x-mathcad",
		registered: false,
		extensions: []string{
			"mcd",
		},
	},
	{
		name:       "application/x-meme",
		registered: false,
		extensions: []string{
			"mm",
		},
	},
	{
		name:       "application/x-midi",
		registered: false,
		extensions: []string{
			"mid", "midi",
		},
	},
	{
		name:       "application/x-mie",
		registered: false,
		extensions: []string{
			"mie",
		},
	},
	{
		name:       "application/x-mif",
		registered: false,
		extensions: []string{
			"mif",
		},
	},
	{
		name:       "application/x-mix-transfer",
		registered: false,
		extensions: []string{
			"nix",
		},
	},
	{
		name:       "application/x-mobipocket-ebook",
		registered: false,
		extensions: []string{
			"prc", "mobi",
		},
	},
	{
		name:       "application/x-mpegURL",
		registered: false,
		extensions: []string{
			"m3u8",
		},
	},
	{
		name:       "application/x-mplayer2",
		registered: false,
		extensions: []string{
			"asx",
		},
	},
	{
		name:       "application/x-ms-application",
		registered: false,
		extensions: []string{
			"application",
		},
	},
	{
		name:       "application/x-ms-shortcut",
		registered: false,
		extensions: []string{
			"lnk",
		},
	},
	{
		name:       "application/x-ms-wmd",
		registered: false,
		extensions: []string{
			"wmd",
		},
	},
	{
		name:       "application/x-ms-wmz",
		registered: false,
		extensions: []string{
			"wmz",
		},
	},
	{
		name:       "application/x-ms-xbap",
		registered: false,
		extensions: []string{
			"xbap",
		},
	},
	{
		name:       "application/x-msaccess",
		registered: false,
		extensions: []string{
			"mdb",
		},
	},
	{
		name:       "application/x-msbinder",
		registered: false,
		extensions: []string{
			"obd",
		},
	},
	{
		name:       "application/x-mscardfile",
		registered: false,
		extensions: []string{
			"crd",
		},
	},
	{
		name:       "application/x-msclip",
		registered: false,
		extensions: []string{
			"clp",
		},
	},
	{
		name:       "application/x-msdos-program",
		registered: false,
		extensions: []string{
			"com", "exe", "bat", "dll",
		},
	},
	{
		name:       "application/x-msdownload",
		registered: false,
		extensions: []string{
			"exe", "dll", "com", "bat", "msi",
		},
	},
	{
		name:       "application/x-msexcel",
		registered: false,
		extensions: []string{
			"xla", "xls", "xlw",
		},
	},
	{
		name:       "application/x-msi",
		registered: false,
		extensions: []string{
			"msi",
		},
	},
	{
		name:       "application/x-msmediaview",
		registered: false,
		extensions: []string{
			"mvb", "m13", "m14",
		},
	},
	{
		name:       "application/x-msmetafile",
		registered: false,
		extensions: []string{
			"wmf", "wmz", "emf", "emz",
		},
	},
	{
		name:       "application/x-msmoney",
		registered: false,
		extensions: []string{
			"mny",
		},
	},
	{
		name:       "application/x-mspowerpoint",
		registered: false,
		extensions: []string{
			"ppt",
		},
	},
	{
		name:       "application/x-mspublisher",
		registered: false,
		extensions: []string{
			"pub",
		},
	},
	{
		name:       "application/x-msschedule",
		registered: false,
		extensions: []string{
			"scd",
		},
	},
	{
		name:       "application/x-msterminal",
		registered: false,
		extensions: []string{
			"trm",
		},
	},
	{
		name:       "application/x-mswrite",
		registered: false,
		extensions: []string{
			"wri",
		},
	},
	{
		name:       "application/x-navi-animation",
		registered: false,
		extensions: []string{
			"ani",
		},
	},
	{
		name:       "application/x-navidoc",
		registered: false,
		extensions: []string{
			"nvd",
		},
	},
	{
		name:       "application/x-navimap",
		registered: false,
		extensions: []string{
			"map",
		},
	},
	{
		name:       "application/x-navistyle",
		registered: false,
		extensions: []string{
			"stl",
		},
	},
	{
		name:       "application/x-netcdf",
		registered: false,
		extensions: []string{
			"nc", "cdf",
		},
	},
	{
		name:       "application/x-newton-compatible-pkg",
		registered: false,
		extensions: []string{
			"pkg",
		},
	},
	{
		name:       "application/x-nokia-9000-communicator-add-on-software",
		registered: false,
		extensions: []string{
			"aos",
		},
	},
	{
		name:       "application/x-ns-proxy-autoconfig",
		registered: false,
		extensions: []string{
			"pac",
		},
	},
	{
		name:       "application/x-nwc",
		registered: false,
		extensions: []string{
			"nwc",
		},
	},
	{
		name:       "application/x-nzb",
		registered: false,
		extensions: []string{
			"nzb",
		},
	},
	{
		name:       "application/x-object",
		registered: false,
		extensions: []string{
			"o",
		},
	},
	{
		name:       "application/x-omc",
		registered: false,
		extensions: []string{
			"omc",
		},
	},
	{
		name:       "application/x-omcdatamaker",
		registered: false,
		extensions: []string{
			"omcd",
		},
	},
	{
		name:       "application/x-omcregerator",
		registered: false,
		extensions: []string{
			"omcr",
		},
	},
	{
		name:       "application/x-oz-application",
		registered: false,
		extensions: []string{
			"oza",
		},
	},
	{
		name:       "application/x-pagemaker",
		registered: false,
		extensions: []string{
			"pm4", "pm5",
		},
	},
	{
		name:       "application/x-pcl",
		registered: false,
		extensions: []string{
			"pcl",
		},
	},
	{
		name:       "application/x-perfmon",
		registered: false,
		extensions: []string{
			"pma", "pmc", "pml", "pmr", "pmw",
		},
	},
	{
		name:       "application/x-pixclscript",
		registered: false,
		extensions: []string{
			"plx",
		},
	},
	{
		name:       "application/x-pkcs10",
		registered: false,
		extensions: []string{
			"p10",
		},
	},
	{
		name:       "application/x-pkcs12",
		registered: false,
		extensions: []string{
			"p12", "pfx",
		},
	},
	{
		name:       "application/x-pkcs7-certificates",
		registered: false,
		extensions: []string{
			"p7b", "spc",
		},
	},
	{
		name:       "application/x-pkcs7-certreqresp",
		registered: false,
		extensions: []string{
			"p7r",
		},
	},
	{
		name:       "application/x-pkcs7-crl",
		registered: false,
		extensions: []string{
			"crl",
		},
	},
	{
		name:       "application/x-pkcs7-mime",
		registered: false,
		extensions: []string{
			"p7c", "p7m",
		},
	},
	{
		name:       "application/x-pkcs7-signature",
		registered: false,
		extensions: []string{
			"p7a", "p7s",
		},
	},
	{
		name:       "application/x-pki-message",
		registered: true,
	},
	{
		name:       "application/x-pointplus",
		registered: false,
		extensions: []string{
			"css",
		},
	},
	{
		name:       "application/x-portable-anymap",
		registered: false,
		extensions: []string{
			"pnm",
		},
	},
	{
		name:       "application/x-project",
		registered: false,
		extensions: []string{
			"mpc", "mpt", "mpv", "mpx",
		},
	},
	{
		name:       "application/x-python-code",
		registered: false,
		extensions: []string{
			"pyc", "pyo",
		},
	},
	{
		name:       "application/x-qpro",
		registered: false,
		extensions: []string{
			"wb1",
		},
	},
	{
		name:       "application/x-quicktimeplayer",
		registered: false,
		extensions: []string{
			"qtl",
		},
	},
	{
		name:       "application/x-rar-compressed",
		registered: false,
		extensions: []string{
			"rar",
		},
	},
	{
		name:       "application/x-redhat-package-manager",
		registered: false,
		extensions: []string{
			"rpm",
		},
	},
	{
		name:       "application/x-research-info-systems",
		registered: false,
		extensions: []string{
			"ris",
		},
	},
	{
		name:       "application/x-rpm",
		registered: false,
		extensions: []string{
			"rpm",
		},
	},
	{
		name:       "application/x-rtf",
		registered: false,
		extensions: []string{
			"rtf",
		},
	},
	{
		name:       "application/x-rx",
		registered: false,
	},
	{
		name:       "application/x-sdp",
		registered: false,
		extensions: []string{
			"sdp",
		},
	},
	{
		name:       "application/x-sea",
		registered: false,
		extensions: []string{
			"sea",
		},
	},
	{
		name:       "application/x-seelogo",
		registered: false,
		extensions: []string{
			"sl",
		},
	},
	{
		name:       "application/x-sh",
		registered: false,
		extensions: []string{
			"sh",
		},
	},
	{
		name:       "application/x-shar",
		registered: false,
		extensions: []string{
			"shar", "sh",
		},
	},
	{
		name:       "application/x-shellscript",
		registered: false,
	},
	{
		name:       "application/x-shockwave-flash",
		registered: false,
		extensions: []string{
			"swf", "swfl",
		},
	},
	{
		name:       "application/x-silverlight-app",
		registered: false,
		extensions: []string{
			"xap",
		},
	},
	{
		name:       "application/x-sit",
		registered: false,
		extensions: []string{
			"sit",
		},
	},
	{
		name:       "application/x-sprite",
		registered: false,
		extensions: []string{
			"spr", "sprite",
		},
	},
	{
		name:       "application/x-sql",
		registered: false,
		extensions: []string{
			"sql",
		},
	},
	{
		name:       "application/x-stuffit",
		registered: false,
		extensions: []string{
			"sit",
		},
	},
	{
		name:       "application/x-stuffitx",
		registered: false,
		extensions: []string{
			"sitx",
		},
	},
	{
		name:       "application/x-subrip",
		registered: false,
		extensions: []string{
			"srt",
		},
	},
	{
		name:       "application/x-sv4cpio",
		registered: false,
		extensions: []string{
			"sv4cpio",
		},
	},
	{
		name:       "application/x-sv4crc",
		registered: false,
		extensions: []string{
			"sv4crc",
		},
	},
	{
		name:       "application/x-t3vm-image",
		registered: false,
		extensions: []string{
			"t3",
		},
	},
	{
		name:       "application/x-tads",
		registered: false,
		extensions: []string{
			"gam",
		},
	},
	{
		name:       "application/x-tar",
		registered: false,
		extensions: []string{
			"tar",
		},
	},
	{
		name:       "application/x-tbook",
		registered: false,
		extensions: []string{
			"sbk", "tbk",
		},
	},
	{
		name:       "application/x-tcl",
		registered: false,
		extensions: []string{
			"tcl",
		},
	},
	{
		name:       "application/x-tex",
		registered: false,
		extensions: []string{
			"tex",
		},
	},
	{
		name:       "application/x-tex-gf",
		registered: false,
		extensions: []string{
			"gf",
		},
	},
	{
		name:       "application/x-tex-pk",
		registered: false,
		extensions: []string{
			"pk",
		},
	},
	{
		name:       "application/x-tex-tfm",
		registered: false,
		extensions: []string{
			"tfm",
		},
	},
	{
		name:       "application/x-texinfo",
		registered: false,
		extensions: []string{
			"texinfo", "texi",
		},
	},
	{
		name:       "application/x-tgif",
		registered: false,
		extensions: []string{
			"obj",
		},
	},
	{
		name:       "application/x-trash",
		registered: false,
		extensions: []string{
			"~", "%", "bak", "old", "sik",
		},
	},
	{
		name:       "application/x-troff",
		registered: false,
		extensions: []string{
			"roff", "t", "tr",
		},
	},
	{
		name:       "application/x-troff-man",
		registered: false,
		extensions: []string{
			"man",
		},
	},
	{
		name:       "application/x-troff-me",
		registered: false,
		extensions: []string{
			"me",
		},
	},
	{
		name:       "application/x-troff-ms",
		registered: false,
		extensions: []string{
			"ms",
		},
	},
	{
		name:       "application/x-troff-msvideo",
		registered: false,
		extensions: []string{
			"avi",
		},
	},
	{
		name:       "application/x-ustar",
		registered: false,
		extensions: []string{
			"ustar",
		},
	},
	{
		name:       "application/x-videolan",
		registered: false,
	},
	{
		name:       "application/x-visio",
		registered: false,
		extensions: []string{
			"vsd", "vst", "vsw",
		},
	},
	{
		name:       "application/x-vnd.audioexplosion.mzz",
		registered: false,
		extensions: []string{
			"mzz",
		},
	},
	{
		name:       "application/x-vnd.ls-xpix",
		registered: false,
		extensions: []string{
			"xpix",
		},
	},
	{
		name:       "application/x-vrml",
		registered: false,
		extensions: []string{
			"vrml",
		},
	},
	{
		name:       "application/x-wais-source",
		registered: false,
		extensions: []string{
			"src", "wsrc",
		},
	},
	{
		name:       "application/x-web-app-manifest+json",
		format:     "application/json",
		registered: false,
		extensions: []string{
			"webapp",
		},
	},
	{
		name:       "application/x-wingz",
		registered: false,
		extensions: []string{
			"wz",
		},
	},
	{
		name:       "application/x-winhelp",
		registered: false,
		extensions: []string{
			"hlp",
		},
	},
	{
		name:       "application/x-wintalk",
		registered: false,
		extensions: []string{
			"wtk",
		},
	},
	{
		name:       "application/x-world",
		registered: false,
		extensions: []string{
			"svr", "wrl",
		},
	},
	{
		name:       "application/x-wpwin",
		registered: false,
		extensions: []string{
			"wpd",
		},
	},
	{
		name:       "application/x-wri",
		registered: false,
		extensions: []string{
			"wri",
		},
	},
	{
		name:       "application/x-www-form-urlencoded",
		registered: true,
	},
	{
		name:       "application/x-x509-ca-cert",
		registered: true,
		extensions: []string{
			"der", "cer", "crt",
		},
	},
	{
		name:       "application/x-x509-ca-ra-cert",
		registered: true,
	},
	{
		name:       "application/x-x509-next-ca-cert",
		registered: true,
	},
	{
		name:       "application/x-x509-user-cert",
		registered: false,
		extensions: []string{
			"crt",
		},
	},
	{
		name:       "application/x-xcf",
		registered: false,
		extensions: []string{
			"xcf",
		},
	},
	{
		name:       "application/x-xfig",
		registered: false,
		extensions: []string{
			"fig",
		},
	},
	{
		name:       "application/x-xliff+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"xlf",
		},
	},
	{
		name:       "application/x-xpinstall",
		registered: false,
		extensions: []string{
			"xpi",
		},
	},
	{
		name:       "application/x-xz",
		registered: false,
		extensions: []string{
			"xz",
		},
	},
	{
		name:       "application/x-zip-compressed",
		registered: false,
		extensions: []string{
			"zip",
		},
	},
	{
		name:       "application/x-zmachine",
		registered: false,
		extensions: []string{
			"z1", "z2", "z3", "z4", "z5", "z6", "z7", "z8",
		},
	},
	{
		name:       "application/x400-bp",
		registered: true,
	},
	{
		name:       "application/xacml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xaml+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"xaml",
		},
	},
	{
		name:       "application/xcap-att+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xcap-caps+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xcap-diff+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xdf",
		},
	},
	{
		name:       "application/xcap-el+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xcap-error+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xcap-ns+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xcon-conference-info+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xcon-conference-info-diff+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xenc+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xenc",
		},
	},
	{
		name:       "application/xfdf",
		registered: true,
	},
	{
		name:       "application/xhtml+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xhtml", "xht",
		},
	},
	{
		name:       "application/xhtml-voice+xml",
		format:     "text/xml",
		registered: false,
	},
	{
		name:       "application/xliff+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xml",
		registered: true,
		extensions: []string{
			"xml", "xsl", "xpdl",
		},
	},
	{
		name:       "application/xml-dtd",
		registered: true,
		extensions: []string{
			"dtd",
		},
	},
	{
		name:       "application/xml-external-parsed-entity",
		registered: true,
	},
	{
		name:       "application/xml-patch+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xmpp+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/xop+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xop",
		},
	},
	{
		name:       "application/xproc+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"xpl",
		},
	},
	{
		name:       "application/xslt+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"xslt",
		},
	},
	{
		name:       "application/xspf+xml",
		format:     "text/xml",
		registered: false,
		extensions: []string{
			"xspf",
		},
	},
	{
		name:       "application/xv+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"mxml", "xhvml", "xvml", "xvm",
		},
	},
	{
		name:       "application/yang",
		registered: true,
		extensions: []string{
			"yang",
		},
	},
	{
		name:       "application/yang-data+cbor",
		format:     "application/cbor",
		registered: true,
	},
	{
		name:       "application/yang-data+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/yang-data+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/yang-patch+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "application/yang-patch+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "application/yin+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"yin",
		},
	},
	{
		name:       "application/ynd.ms-pkipko",
		registered: false,
		extensions: []string{
			"pko",
		},
	},
	{
		name:       "application/zip",
		registered: true,
		extensions: []string{
			"zip",
		},
	},
	{
		name:       "application/zlib",
		registered: true,
	},
	{
		name:       "application/zstd",
		registered: true,
	},
	{
		name:       "audio/1d-interleaved-parityfec",
		registered: true,
	},
	{
		name:       "audio/32kadpcm",
		registered: true,
	},
	{
		name:       "audio/3gpp",
		registered: true,
	},
	{
		name:       "audio/3gpp2",
		registered: true,
	},
	{
		name:       "audio/ATRAC-X",
		registered: true,
	},
	{
		name:       "audio/ATRAC3",
		registered: true,
	},
	{
		name:       "audio/BV32",
		registered: true,
	},
	{
		name:       "audio/CN",
		registered: true,
	},
	{
		name:       "audio/DV",
		registered: true,
	},
	{
		name:       "audio/EVRC",
		registered: true,
	},
	{
		name:       "audio/EVRC1",
		registered: true,
	},
	{
		name:       "audio/EVRCB",
		registered: true,
	},
	{
		name:       "audio/EVRCB0",
		registered: true,
	},
	{
		name:       "audio/EVRCB1",
		registered: true,
	},
	{
		name:       "audio/EVRCNW",
		registered: true,
	},
	{
		name:       "audio/EVRCNW0",
		registered: true,
	},
	{
		name:       "audio/EVRCNW1",
		registered: true,
	},
	{
		name:       "audio/EVRCWB",
		registered: true,
	},
	{
		name:       "audio/EVRCWB0",
		registered: true,
	},
	{
		name:       "audio/EVRCWB1",
		registered: true,
	},
	{
		name:       "audio/EVS",
		registered: true,
	},
	{
		name:       "audio/G711-0",
		registered: true,
	},
	{
		name:       "audio/G722",
		registered: true,
	},
	{
		name:       "audio/G7221",
		registered: true,
	},
	{
		name:       "audio/G726-16",
		registered: true,
	},
	{
		name:       "audio/G729",
		registered: true,
	},
	{
		name:       "audio/G7291",
		registered: true,
	},
	{
		name:       "audio/GSM",
		registered: true,
	},
	{
		name:       "audio/L24",
		registered: true,
	},
	{
		name:       "audio/L8",
		registered: true,
	},
	{
		name:       "audio/MELP",
		registered: true,
	},
	{
		name:       "audio/MELP1200",
		registered: true,
	},
	{
		name:       "audio/MELP2400",
		registered: true,
	},
	{
		name:       "audio/MELP600",
		registered: true,
	},
	{
		name:       "audio/MP4A-LATM",
		registered: true,
	},
	{
		name:       "audio/MPA",
		registered: true,
	},
	{
		name:       "audio/PCMA-WB",
		registered: true,
	},
	{
		name:       "audio/PCMU-WB",
		registered: true,
	},
	{
		name:       "audio/QCELP",
		registered: true,
	},
	{
		name:       "audio/SMV",
		registered: true,
	},
	{
		name:       "audio/SMV-QCP",
		registered: true,
	},
	{
		name:       "audio/SMV0",
		registered: true,
	},
	{
		name:       "audio/TETRA_ACELP",
		registered: true,
	},
	{
		name:       "audio/TETRA_ACELP_BB",
		registered: true,
	},
	{
		name:       "audio/TSVCIS",
		registered: true,
	},
	{
		name:       "audio/UEMCLIP",
		registered: true,
	},
	{
		name:       "audio/VMR-WB",
		registered: true,
	},
	{
		name:       "audio/aac",
		registered: true,
	},
	{
		name:       "audio/ac3",
		registered: true,
	},
	{
		name:       "audio/adpcm",
		registered: false,
		extensions: []string{
			"adp",
		},
	},
	{
		name:       "audio/aiff",
		registered: false,
		extensions: []string{
			"aif", "aifc", "aiff",
		},
	},
	{
		name:       "audio/amr",
		registered: false,
	},
	{
		name:       "audio/amr-wb",
		registered: false,
	},
	{
		name:       "audio/amr-wb+",
		registered: true,
	},
	{
		name:       "audio/aptx",
		registered: true,
	},
	{
		name:       "audio/asc",
		registered: true,
	},
	{
		name:       "audio/atrac-advanced-lossless",
		registered: false,
	},
	{
		name:       "audio/basic",
		registered: true,
		extensions: []string{
			"au", "snd",
		},
	},
	{
		name:       "audio/bv16",
		registered: false,
	},
	{
		name:       "audio/clearmode",
		registered: true,
	},
	{
		name:       "audio/dat12",
		registered: false,
	},
	{
		name:       "audio/dls",
		registered: true,
	},
	{
		name:       "audio/dsr-es201108",
		registered: true,
	},
	{
		name:       "audio/dsr-es202050",
		registered: true,
	},
	{
		name:       "audio/dsr-es202211",
		registered: true,
	},
	{
		name:       "audio/dsr-es202212",
		registered: true,
	},
	{
		name:       "audio/dvi4",
		registered: false,
	},
	{
		name:       "audio/eac3",
		registered: true,
	},
	{
		name:       "audio/encaprtp",
		registered: true,
	},
	{
		name:       "audio/evrc-qcp",
		registered: false,
	},
	{
		name:       "audio/evrc0",
		registered: false,
	},
	{
		name:       "audio/example",
		registered: true,
	},
	{
		name:       "audio/flac",
		registered: false,
		extensions: []string{
			"flac",
		},
	},
	{
		name:       "audio/flexfec",
		registered: true,
	},
	{
		name:       "audio/fwdred",
		registered: true,
	},
	{
		name:       "audio/g.722.1",
		registered: false,
	},
	{
		name:       "audio/g719",
		registered: false,
	},
	{
		name:       "audio/g723",
		registered: false,
	},
	{
		name:       "audio/g726-24",
		registered: false,
	},
	{
		name:       "audio/g726-32",
		registered: false,
	},
	{
		name:       "audio/g726-40",
		registered: false,
	},
	{
		name:       "audio/g728",
		registered: false,
	},
	{
		name:       "audio/g729d",
		registered: false,
	},
	{
		name:       "audio/g729e",
		registered: false,
	},
	{
		name:       "audio/gsm-efr",
		registered: false,
	},
	{
		name:       "audio/gsm-hr-08",
		registered: false,
	},
	{
		name:       "audio/iLBC",
		registered: true,
	},
	{
		name:       "audio/ip-mr_v2.5",
		registered: true,
	},
	{
		name:       "audio/isac",
		registered: false,
	},
	{
		name:       "audio/it",
		registered: false,
		extensions: []string{
			"it",
		},
	},
	{
		name:       "audio/l16",
		registered: false,
	},
	{
		name:       "audio/l20",
		registered: false,
	},
	{
		name:       "audio/lpc",
		registered: false,
	},
	{
		name:       "audio/make",
		registered: false,
		extensions: []string{
			"funk", "my", "pfunk",
		},
	},
	{
		name:       "audio/make.my.funk",
		registered: false,
		extensions: []string{
			"pfunk",
		},
	},
	{
		name:       "audio/mhas",
		registered: true,
	},
	{
		name:       "audio/mid",
		registered: false,
		extensions: []string{
			"rmi", "mid",
		},
	},
	{
		name:       "audio/midi",
		registered: false,
		extensions: []string{
			"mid", "midi", "kar", "rmi",
		},
	},
	{
		name:       "audio/mobile-xmf",
		registered: true,
	},
	{
		name:       "audio/mod",
		registered: false,
		extensions: []string{
			"mod",
		},
	},
	{
		name:       "audio/mp4",
		registered: true,
		extensions: []string{
			"mp4a", "m4a",
		},
	},
	{
		name:       "audio/mpa-robust",
		registered: true,
	},
	{
		name:       "audio/mpeg",
		registered: true,
		extensions: []string{
			"mpga", "mp2", "mp2a", "mp3", "m2a", "mpa", "mpg", "m3a", "mpega", "m4a",
		},
	},
	{
		name:       "audio/mpeg3",
		registered: false,
		extensions: []string{
			"mp3",
		},
	},
	{
		name:       "audio/mpeg4-generic",
		registered: true,
	},
	{
		name:       "audio/mpegurl",
		registered: false,
		extensions: []string{
			"m3u",
		},
	},
	{
		name:       "audio/musepack",
		registered: false,
	},
	{
		name:       "audio/nspaudio",
		registered: false,
		extensions: []string{
			"la", "lma",
		},
	},
	{
		name:       "audio/ogg",
		registered: true,
		extensions: []string{
			"oga", "ogg", "spx",
		},
	},
	{
		name:       "audio/opus",
		registered: true,
	},
	{
		name:       "audio/parityfec",
		registered: true,
	},
	{
		name:       "audio/pcma",
		registered: false,
	},
	{
		name:       "audio/pcmu",
		registered: false,
	},
	{
		name:       "audio/prs.sid",
		registered: true,
		extensions: []string{
			"sid",
		},
	},
	{
		name:       "audio/raptorfec",
		registered: true,
	},
	{
		name:       "audio/red",
		registered: false,
	},
	{
		name:       "audio/rtp-enc-aescm128",
		registered: true,
	},
	{
		name:       "audio/rtp-midi",
		registered: true,
	},
	{
		name:       "audio/rtploopback",
		registered: true,
	},
	{
		name:       "audio/rtx",
		registered: true,
	},
	{
		name:       "audio/s3m",
		registered: false,
		extensions: []string{
			"s3m",
		},
	},
	{
		name:       "audio/scip",
		registered: true,
	},
	{
		name:       "audio/silk",
		registered: false,
		extensions: []string{
			"sil",
		},
	},
	{
		name:       "audio/sofa",
		registered: true,
	},
	{
		name:       "audio/sp-midi",
		registered: true,
	},
	{
		name:       "audio/speex",
		registered: true,
	},
	{
		name:       "audio/t140c",
		registered: true,
	},
	{
		name:       "audio/t38",
		registered: true,
	},
	{
		name:       "audio/telephone-event",
		registered: true,
	},
	{
		name:       "audio/tone",
		registered: true,
	},
	{
		name:       "audio/tsp-audio",
		registered: false,
		extensions: []string{
			"tsi",
		},
	},
	{
		name:       "audio/tsplayer",
		registered: false,
		extensions: []string{
			"tsp",
		},
	},
	{
		name:       "audio/ulpfec",
		registered: true,
	},
	{
		name:       "audio/usac",
		registered: true,
	},
	{
		name:       "audio/vdvi",
		registered: false,
	},
	{
		name:       "audio/vnd.3gpp.iufp",
		registered: true,
	},
	{
		name:       "audio/vnd.4SB",
		registered: true,
	},
	{
		name:       "audio/vnd.CELP",
		registered: true,
	},
	{
		name:       "audio/vnd.audiokoz",
		registered: true,
	},
	{
		name:       "audio/vnd.cisco.nse",
		registered: true,
	},
	{
		name:       "audio/vnd.cmles.radio-events",
		registered: true,
	},
	{
		name:       "audio/vnd.cns.anp1",
		registered: true,
	},
	{
		name:       "audio/vnd.cns.inf1",
		registered: true,
	},
	{
		name:       "audio/vnd.dece.audio",
		registered: true,
		extensions: []string{
			"uva", "uvva",
		},
	},
	{
		name:       "audio/vnd.digital-winds",
		registered: true,
		extensions: []string{
			"eol",
		},
	},
	{
		name:       "audio/vnd.dlna.adts",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.heaac.1",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.heaac.2",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.mlp",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.mps",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.pl2",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.pl2x",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.pl2z",
		registered: true,
	},
	{
		name:       "audio/vnd.dolby.pulse.1",
		registered: true,
	},
	{
		name:       "audio/vnd.dra",
		registered: true,
		extensions: []string{
			"dra",
		},
	},
	{
		name:       "audio/vnd.dts",
		registered: true,
		extensions: []string{
			"dts",
		},
	},
	{
		name:       "audio/vnd.dts.hd",
		registered: true,
		extensions: []string{
			"dtshd",
		},
	},
	{
		name:       "audio/vnd.dts.uhd",
		registered: true,
	},
	{
		name:       "audio/vnd.dvb.file",
		registered: true,
	},
	{
		name:       "audio/vnd.everad.plj",
		registered: true,
	},
	{
		name:       "audio/vnd.hns.audio",
		registered: true,
	},
	{
		name:       "audio/vnd.lucent.voice",
		registered: true,
		extensions: []string{
			"lvp",
		},
	},
	{
		name:       "audio/vnd.ms-playready.media.pya",
		registered: true,
		extensions: []string{
			"pya",
		},
	},
	{
		name:       "audio/vnd.nokia.mobile-xmf",
		registered: true,
	},
	{
		name:       "audio/vnd.nortel.vbk",
		registered: true,
	},
	{
		name:       "audio/vnd.nuera.ecelp4800",
		registered: true,
		extensions: []string{
			"ecelp4800",
		},
	},
	{
		name:       "audio/vnd.nuera.ecelp7470",
		registered: true,
		extensions: []string{
			"ecelp7470",
		},
	},
	{
		name:       "audio/vnd.nuera.ecelp9600",
		registered: true,
		extensions: []string{
			"ecelp9600",
		},
	},
	{
		name:       "audio/vnd.octel.sbc",
		registered: true,
	},
	{
		name:       "audio/vnd.presonus.multitrack",
		registered: true,
	},
	{
		name:       "audio/vnd.qcelp",
		registered: true,
		extensions: []string{
			"qcp",
		},
	},
	{
		name:       "audio/vnd.rhetorex.32kadpcm",
		registered: true,
	},
	{
		name:       "audio/vnd.rip",
		registered: true,
		extensions: []string{
			"rip",
		},
	},
	{
		name:       "audio/vnd.sealedmedia.softseal.mpeg",
		registered: true,
	},
	{
		name:       "audio/vnd.vmx.cvsd",
		registered: true,
	},
	{
		name:       "audio/voc",
		registered: false,
		extensions: []string{
			"voc",
		},
	},
	{
		name:       "audio/vorbis",
		registered: true,
	},
	{
		name:       "audio/vorbis-config",
		registered: true,
	},
	{
		name:       "audio/voxware",
		registered: false,
		extensions: []string{
			"vox",
		},
	},
	{
		name:       "audio/wav",
		registered: false,
		extensions: []string{
			"wav",
		},
	},
	{
		name:       "audio/webm",
		registered: false,
		extensions: []string{
			"weba",
		},
	},
	{
		name:       "audio/x-aac",
		registered: false,
		extensions: []string{
			"aac",
		},
	},
	{
		name:       "audio/x-adpcm",
		registered: false,
		extensions: []string{
			"snd",
		},
	},
	{
		name:       "audio/x-aiff",
		registered: false,
		extensions: []string{
			"aif", "aiff", "aifc",
		},
	},
	{
		name:       "audio/x-au",
		registered: false,
		extensions: []string{
			"au",
		},
	},
	{
		name:       "audio/x-caf",
		registered: false,
		extensions: []string{
			"caf",
		},
	},
	{
		name:       "audio/x-flac",
		registered: false,
		extensions: []string{
			"flac",
		},
	},
	{
		name:       "audio/x-gsm",
		registered: false,
		extensions: []string{
			"gsd", "gsm",
		},
	},
	{
		name:       "audio/x-jam",
		registered: false,
		extensions: []string{
			"jam",
		},
	},
	{
		name:       "audio/x-liveaudio",
		registered: false,
		extensions: []string{
			"lam",
		},
	},
	{
		name:       "audio/x-matroska",
		registered: false,
		extensions: []string{
			"mka",
		},
	},
	{
		name:       "audio/x-mid",
		registered: false,
		extensions: []string{
			"mid", "midi",
		},
	},
	{
		name:       "audio/x-midi",
		registered: false,
		extensions: []string{
			"mid", "midi",
		},
	},
	{
		name:       "audio/x-mod",
		registered: false,
		extensions: []string{
			"mod",
		},
	},
	{
		name:       "audio/x-mpeg",
		registered: false,
		extensions: []string{
			"mp2",
		},
	},
	{
		name:       "audio/x-mpeg-3",
		registered: false,
		extensions: []string{
			"mp3",
		},
	},
	{
		name:       "audio/x-mpegurl",
		registered: false,
		extensions: []string{
			"m3u",
		},
	},
	{
		name:       "audio/x-mpequrl",
		registered: false,
		extensions: []string{
			"m3u",
		},
	},
	{
		name:       "audio/x-ms-wax",
		registered: false,
		extensions: []string{
			"wax",
		},
	},
	{
		name:       "audio/x-ms-wma",
		registered: false,
		extensions: []string{
			"wma",
		},
	},
	{
		name:       "audio/x-nspaudio",
		registered: false,
		extensions: []string{
			"la", "lma",
		},
	},
	{
		name:       "audio/x-pn-realaudio",
		registered: false,
		extensions: []string{
			"ram", "ra", "rm", "rmm", "rmp",
		},
	},
	{
		name:       "audio/x-pn-realaudio-plugin",
		registered: false,
		extensions: []string{
			"rmp", "ra", "rpm",
		},
	},
	{
		name:       "audio/x-psid",
		registered: false,
		extensions: []string{
			"sid",
		},
	},
	{
		name:       "audio/x-realaudio",
		registered: false,
		extensions: []string{
			"ra",
		},
	},
	{
		name:       "audio/x-scpls",
		registered: false,
		extensions: []string{
			"pls",
		},
	},
	{
		name:       "audio/x-sd2",
		registered: false,
		extensions: []string{
			"sd2",
		},
	},
	{
		name:       "audio/x-tta",
		registered: false,
	},
	{
		name:       "audio/x-twinvq",
		registered: false,
		extensions: []string{
			"vqf",
		},
	},
	{
		name:       "audio/x-twinvq-plugin",
		registered: false,
		extensions: []string{
			"vqe", "vql",
		},
	},
	{
		name:       "audio/x-vnd.audioexplosion.mjuicemediafile",
		registered: false,
		extensions: []string{
			"mjf",
		},
	},
	{
		name:       "audio/x-voc",
		registered: false,
		extensions: []string{
			"voc",
		},
	},
	{
		name:       "audio/x-wav",
		registered: false,
		extensions: []string{
			"wav",
		},
	},
	{
		name:       "audio/xm",
		registered: false,
		extensions: []string{
			"xm",
		},
	},
	{
		name:       "chemical/x-cdx",
		registered: false,
		extensions: []string{
			"cdx",
		},
	},
	{
		name:       "chemical/x-cif",
		registered: false,
		extensions: []string{
			"cif",
		},
	},
	{
		name:       "chemical/x-cmdf",
		registered: false,
		extensions: []string{
			"cmdf",
		},
	},
	{
		name:       "chemical/x-cml",
		registered: false,
		extensions: []string{
			"cml",
		},
	},
	{
		name:       "chemical/x-csml",
		registered: false,
		extensions: []string{
			"csml",
		},
	},
	{
		name:       "chemical/x-pdb",
		registered: false,
		extensions: []string{
			"pdb", "xyz",
		},
	},
	{
		name:       "chemical/x-xyz",
		registered: false,
		extensions: []string{
			"xyz",
		},
	},
	{
		name:       "conference/x-cooltalk",
		registered: false,
		extensions: []string{
			"ice",
		},
	},
	{
		name:       "content/unknown",
		registered: false,
	},
	{
		name:       "drawing/x-dwf",
		registered: false,
		extensions: []string{
			"dwf",
		},
	},
	{
		name:       "font/collection",
		registered: true,
	},
	{
		name:       "font/opentype",
		registered: false,
		extensions: []string{
			"otf",
		},
	},
	{
		name:       "font/otf",
		registered: true,
	},
	{
		name:       "font/sfnt",
		registered: true,
	},
	{
		name:       "font/ttf",
		registered: true,
	},
	{
		name:       "font/woff",
		registered: true,
	},
	{
		name:       "font/woff2",
		registered: true,
	},
	{
		name:       "image/aces",
		registered: true,
	},
	{
		name:       "image/apng",
		registered: true,
	},
	{
		name:       "image/avci",
		registered: true,
	},
	{
		name:       "image/avcs",
		registered: true,
	},
	{
		name:       "image/avif",
		registered: true,
	},
	{
		name:       "image/bmp",
		registered: true,
		extensions: []string{
			"bmp", "bm",
		},
	},
	{
		name:       "image/cgm",
		registered: true,
		extensions: []string{
			"cgm",
		},
	},
	{
		name:       "image/cis-cod",
		registered: false,
		extensions: []string{
			"cod",
		},
	},
	{
		name:       "image/cmu-raster",
		registered: false,
		extensions: []string{
			"ras", "rast",
		},
	},
	{
		name:       "image/dicom-rle",
		registered: true,
	},
	{
		name:       "image/dpx",
		registered: true,
	},
	{
		name:       "image/emf",
		registered: true,
	},
	{
		name:       "image/example",
		registered: true,
	},
	{
		name:       "image/fif",
		registered: false,
		extensions: []string{
			"fif",
		},
	},
	{
		name:       "image/fits",
		registered: true,
	},
	{
		name:       "image/florian",
		registered: false,
		extensions: []string{
			"flo", "turbot",
		},
	},
	{
		name:       "image/g3fax",
		registered: true,
		extensions: []string{
			"g3",
		},
	},
	{
		name:       "image/gif",
		registered: false,
		extensions: []string{
			"gif",
		},
	},
	{
		name:       "image/heic",
		registered: true,
	},
	{
		name:       "image/heic-sequence",
		registered: true,
	},
	{
		name:       "image/heif",
		registered: true,
	},
	{
		name:       "image/heif-sequence",
		registered: true,
	},
	{
		name:       "image/hej2k",
		registered: true,
	},
	{
		name:       "image/hsj2",
		registered: true,
	},
	{
		name:       "image/ief",
		registered: false,
		extensions: []string{
			"ief", "iefs",
		},
	},
	{
		name:       "image/jls",
		registered: true,
	},
	{
		name:       "image/jp2",
		registered: true,
	},
	{
		name:       "image/jpeg",
		registered: false,
		extensions: []string{
			"jpeg", "jpg", "jfif", "jfif-tbnl", "jpe",
		},
	},
	{
		name:       "image/jph",
		registered: true,
	},
	{
		name:       "image/jphc",
		registered: true,
	},
	{
		name:       "image/jpm",
		registered: true,
	},
	{
		name:       "image/jpx",
		registered: true,
	},
	{
		name:       "image/jutvision",
		registered: false,
		extensions: []string{
			"jut",
		},
	},
	{
		name:       "image/jxr",
		registered: true,
	},
	{
		name:       "image/jxrA",
		registered: true,
	},
	{
		name:       "image/jxrS",
		registered: true,
	},
	{
		name:       "image/jxs",
		registered: true,
	},
	{
		name:       "image/jxsc",
		registered: true,
	},
	{
		name:       "image/jxsi",
		registered: true,
	},
	{
		name:       "image/jxss",
		registered: true,
	},
	{
		name:       "image/ktx",
		registered: true,
		extensions: []string{
			"ktx",
		},
	},
	{
		name:       "image/ktx2",
		registered: true,
	},
	{
		name:       "image/naplps",
		registered: true,
		extensions: []string{
			"nap", "naplps",
		},
	},
	{
		name:       "image/pcx",
		registered: false,
		extensions: []string{
			"pcx",
		},
	},
	{
		name:       "image/pict",
		registered: false,
		extensions: []string{
			"pic", "pict",
		},
	},
	{
		name:       "image/pipeg",
		registered: false,
		extensions: []string{
			"jfif",
		},
	},
	{
		name:       "image/pjpeg",
		registered: false,
		extensions: []string{
			"jfif", "jpe", "jpeg", "jpg",
		},
	},
	{
		name:       "image/png",
		registered: true,
		extensions: []string{
			"png", "x-png",
		},
	},
	{
		name:       "image/prs.btif",
		registered: true,
		extensions: []string{
			"btif",
		},
	},
	{
		name:       "image/prs.pti",
		registered: true,
	},
	{
		name:       "image/pwg-raster",
		registered: true,
	},
	{
		name:       "image/sgi",
		registered: false,
		extensions: []string{
			"sgi",
		},
	},
	{
		name:       "image/svg+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"svg", "svgz",
		},
	},
	{
		name:       "image/t38",
		registered: true,
	},
	{
		name:       "image/tiff",
		registered: true,
		extensions: []string{
			"tiff", "tif",
		},
	},
	{
		name:       "image/tiff-fx",
		registered: true,
	},
	{
		name:       "image/vasa",
		registered: false,
		extensions: []string{
			"mcf",
		},
	},
	{
		name:       "image/vnd.adobe.photoshop",
		registered: true,
		extensions: []string{
			"psd",
		},
	},
	{
		name:       "image/vnd.airzip.accelerator.azv",
		registered: true,
	},
	{
		name:       "image/vnd.cns.inf2",
		registered: true,
	},
	{
		name:       "image/vnd.dece.graphic",
		registered: true,
		extensions: []string{
			"uvi", "uvvi", "uvg", "uvvg",
		},
	},
	{
		name:       "image/vnd.djvu",
		registered: true,
		extensions: []string{
			"djvu", "djv",
		},
	},
	{
		name:       "image/vnd.dvb.subtitle",
		registered: true,
		extensions: []string{
			"sub",
		},
	},
	{
		name:       "image/vnd.dwg",
		registered: true,
		extensions: []string{
			"dwg", "dxf", "svf",
		},
	},
	{
		name:       "image/vnd.dxf",
		registered: true,
		extensions: []string{
			"dxf",
		},
	},
	{
		name:       "image/vnd.fastbidsheet",
		registered: true,
		extensions: []string{
			"fbs",
		},
	},
	{
		name:       "image/vnd.fpx",
		registered: true,
		extensions: []string{
			"fpx", "fpix",
		},
	},
	{
		name:       "image/vnd.fst",
		registered: true,
		extensions: []string{
			"fst",
		},
	},
	{
		name:       "image/vnd.fujixerox.edmics-mmr",
		registered: true,
		extensions: []string{
			"mmr",
		},
	},
	{
		name:       "image/vnd.fujixerox.edmics-rlc",
		registered: true,
		extensions: []string{
			"rlc",
		},
	},
	{
		name:       "image/vnd.globalgraphics.pgb",
		registered: true,
	},
	{
		name:       "image/vnd.microsoft.icon",
		registered: true,
	},
	{
		name:       "image/vnd.mix",
		registered: true,
	},
	{
		name:       "image/vnd.mozilla.apng",
		registered: true,
	},
	{
		name:       "image/vnd.ms-modi",
		registered: true,
		extensions: []string{
			"mdi",
		},
	},
	{
		name:       "image/vnd.ms-photo",
		registered: false,
		extensions: []string{
			"wdp",
		},
	},
	{
		name:       "image/vnd.net-fpx",
		registered: true,
		extensions: []string{
			"npx", "fpx",
		},
	},
	{
		name:       "image/vnd.pco.b16",
		registered: true,
	},
	{
		name:       "image/vnd.radiance",
		registered: true,
	},
	{
		name:       "image/vnd.rn-realflash",
		registered: false,
		extensions: []string{
			"rf",
		},
	},
	{
		name:       "image/vnd.rn-realpix",
		registered: false,
		extensions: []string{
			"rp",
		},
	},
	{
		name:       "image/vnd.sealed.png",
		registered: true,
	},
	{
		name:       "image/vnd.sealedmedia.softseal.gif",
		registered: true,
	},
	{
		name:       "image/vnd.sealedmedia.softseal.jpg",
		registered: true,
	},
	{
		name:       "image/vnd.svf",
		registered: true,
	},
	{
		name:       "image/vnd.tencent.tap",
		registered: true,
	},
	{
		name:       "image/vnd.valve.source.texture",
		registered: true,
	},
	{
		name:       "image/vnd.wap.wbmp",
		registered: true,
		extensions: []string{
			"wbmp",
		},
	},
	{
		name:       "image/vnd.xiff",
		registered: true,
		extensions: []string{
			"xif",
		},
	},
	{
		name:       "image/vnd.zbrush.pcx",
		registered: true,
	},
	{
		name:       "image/webp",
		registered: false,
		extensions: []string{
			"webp",
		},
	},
	{
		name:       "image/wmf",
		registered: true,
	},
	{
		name:       "image/x-3ds",
		registered: false,
		extensions: []string{
			"3ds",
		},
	},
	{
		name:       "image/x-cmu-rast",
		registered: false,
		extensions: []string{
			"ras",
		},
	},
	{
		name:       "image/x-cmu-raster",
		registered: false,
		extensions: []string{
			"ras",
		},
	},
	{
		name:       "image/x-cmx",
		registered: false,
		extensions: []string{
			"cmx",
		},
	},
	{
		name:       "image/x-coreldraw",
		registered: false,
		extensions: []string{
			"cdr",
		},
	},
	{
		name:       "image/x-coreldrawpattern",
		registered: false,
		extensions: []string{
			"pat",
		},
	},
	{
		name:       "image/x-coreldrawtemplate",
		registered: false,
		extensions: []string{
			"cdt",
		},
	},
	{
		name:       "image/x-corelphotopaint",
		registered: false,
		extensions: []string{
			"cpt",
		},
	},
	{
		name:       "image/x-dwg",
		registered: false,
		extensions: []string{
			"dwg", "dxf", "svf",
		},
	},
	{
		name:       "image/x-freehand",
		registered: false,
		extensions: []string{
			"fh", "fhc", "fh4", "fh5", "fh7",
		},
	},
	{
		name:       "image/x-icon",
		registered: false,
		extensions: []string{
			"ico",
		},
	},
	{
		name:       "image/x-jg",
		registered: false,
		extensions: []string{
			"art",
		},
	},
	{
		name:       "image/x-jng",
		registered: false,
		extensions: []string{
			"jng",
		},
	},
	{
		name:       "image/x-jps",
		registered: false,
		extensions: []string{
			"jps",
		},
	},
	{
		name:       "image/x-mrsid-image",
		registered: false,
		extensions: []string{
			"sid",
		},
	},
	{
		name:       "image/x-ms-bmp",
		registered: false,
		extensions: []string{
			"bmp",
		},
	},
	{
		name:       "image/x-niff",
		registered: false,
		extensions: []string{
			"nif", "niff",
		},
	},
	{
		name:       "image/x-pcx",
		registered: false,
		extensions: []string{
			"pcx",
		},
	},
	{
		name:       "image/x-photoshop",
		registered: false,
		extensions: []string{
			"psd",
		},
	},
	{
		name:       "image/x-pict",
		registered: false,
		extensions: []string{
			"pic", "pct",
		},
	},
	{
		name:       "image/x-portable-anymap",
		registered: false,
		extensions: []string{
			"pnm",
		},
	},
	{
		name:       "image/x-portable-bitmap",
		registered: false,
		extensions: []string{
			"pbm",
		},
	},
	{
		name:       "image/x-portable-graymap",
		registered: false,
		extensions: []string{
			"pgm",
		},
	},
	{
		name:       "image/x-portable-greymap",
		registered: false,
		extensions: []string{
			"pgm",
		},
	},
	{
		name:       "image/x-portable-pixmap",
		registered: false,
		extensions: []string{
			"ppm",
		},
	},
	{
		name:       "image/x-quicktime",
		registered: false,
		extensions: []string{
			"qif", "qti", "qtif",
		},
	},
	{
		name:       "image/x-rgb",
		registered: false,
		extensions: []string{
			"rgb",
		},
	},
	{
		name:       "image/x-tga",
		registered: false,
		extensions: []string{
			"tga",
		},
	},
	{
		name:       "image/x-tiff",
		registered: false,
		extensions: []string{
			"tif", "tiff",
		},
	},
	{
		name:       "image/x-windows-bmp",
		registered: false,
		extensions: []string{
			"bmp",
		},
	},
	{
		name:       "image/x-xbitmap",
		registered: false,
		extensions: []string{
			"xbm", "xpm",
		},
	},
	{
		name:       "image/x-xbm",
		registered: false,
		extensions: []string{
			"xbm",
		},
	},
	{
		name:       "image/x-xpixmap",
		registered: false,
		extensions: []string{
			"xpm", "pm",
		},
	},
	{
		name:       "image/x-xwd",
		registered: false,
		extensions: []string{
			"xwd",
		},
	},
	{
		name:       "image/x-xwindowdump",
		registered: false,
		extensions: []string{
			"xwd",
		},
	},
	{
		name:       "image/xbm",
		registered: false,
		extensions: []string{
			"xbm",
		},
	},
	{
		name:       "image/xpm",
		registered: false,
		extensions: []string{
			"xpm",
		},
	},
	{
		name:       "inode/blockdevice",
		registered: false,
	},
	{
		name:       "inode/chardevice",
		registered: false,
	},
	{
		name:       "inode/directory",
		registered: false,
	},
	{
		name:       "inode/directory-locked",
		registered: false,
	},
	{
		name:       "inode/fifo",
		registered: false,
	},
	{
		name:       "inode/socket",
		registered: false,
	},
	{
		name:       "message/bhttp",
		registered: true,
	},
	{
		name:       "message/cpim",
		registered: false,
	},
	{
		name:       "message/delivery-status",
		registered: true,
	},
	{
		name:       "message/disposition-notification",
		registered: true,
	},
	{
		name:       "message/example",
		registered: true,
	},
	{
		name:       "message/external-body",
		registered: false,
	},
	{
		name:       "message/feedback-report",
		registered: true,
	},
	{
		name:       "message/global",
		registered: true,
	},
	{
		name:       "message/global-delivery-status",
		registered: true,
	},
	{
		name:       "message/global-disposition-notification",
		registered: true,
	},
	{
		name:       "message/global-headers",
		registered: true,
	},
	{
		name:       "message/http",
		registered: true,
	},
	{
		name:       "message/imdn+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "message/news",
		registered: true,
	},
	{
		name:       "message/partial",
		registered: false,
	},
	{
		name:       "message/rfc822",
		registered: false,
		extensions: []string{
			"eml", "mht", "mhtml", "mime", "nws",
		},
	},
	{
		name:       "message/s-http",
		registered: true,
	},
	{
		name:       "message/sip",
		registered: true,
	},
	{
		name:       "message/sipfrag",
		registered: true,
	},
	{
		name:       "message/tracking-status",
		registered: true,
	},
	{
		name:       "message/vnd.si.simp",
		registered: true,
	},
	{
		name:       "message/vnd.wfa.wsc",
		registered: true,
	},
	{
		name:       "model/3mf",
		registered: true,
	},
	{
		name:       "model/e57",
		registered: true,
	},
	{
		name:       "model/example",
		registered: true,
	},
	{
		name:       "model/gltf+json",
		format:     "application/json",
		registered: true,
	},
	{
		name:       "model/gltf-binary",
		registered: true,
	},
	{
		name:       "model/iges",
		registered: true,
		extensions: []string{
			"igs", "iges",
		},
	},
	{
		name:       "model/mesh",
		registered: false,
		extensions: []string{
			"msh", "mesh", "silo",
		},
	},
	{
		name:       "model/mtl",
		registered: true,
	},
	{
		name:       "model/obj",
		registered: true,
	},
	{
		name:       "model/prc",
		registered: true,
	},
	{
		name:       "model/step",
		registered: true,
	},
	{
		name:       "model/step+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "model/step+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "model/step-xml+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "model/stl",
		registered: true,
	},
	{
		name:       "model/u3d",
		registered: true,
	},
	{
		name:       "model/vnd.collada+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"dae",
		},
	},
	{
		name:       "model/vnd.dwf",
		registered: true,
		extensions: []string{
			"dwf",
		},
	},
	{
		name:       "model/vnd.flatland.3dml",
		registered: true,
	},
	{
		name:       "model/vnd.gdl",
		registered: true,
		extensions: []string{
			"gdl",
		},
	},
	{
		name:       "model/vnd.gs-gdl",
		registered: true,
	},
	{
		name:       "model/vnd.gs.gdl",
		registered: false,
	},
	{
		name:       "model/vnd.gtw",
		registered: true,
		extensions: []string{
			"gtw",
		},
	},
	{
		name:       "model/vnd.moml+xml",
		format:     "text/xml",
		registered: true,
	},
	{
		name:       "model/vnd.mts",
		registered: true,
		extensions: []string{
			"mts",
		},
	},
	{
		name:       "model/vnd.opengex",
		registered: true,
	},
	{
		name:       "model/vnd.parasolid.transmit.binary",
		registered: true,
	},
	{
		name:       "model/vnd.parasolid.transmit.text",
		registered: true,
	},
	{
		name:       "model/vnd.pytha.pyox",
		registered: true,
	},
	{
		name:       "model/vnd.rosette.annotated-data-model",
		registered: true,
	},
	{
		name:       "model/vnd.sap.vds",
		registered: true,
	},
	{
		name:       "model/vnd.usda",
		registered: true,
	},
	{
		name:       "model/vnd.usdz+zip",
		format:     "application/zip",
		registered: true,
	},
	{
		name:       "model/vnd.valve.source.compiled-map",
		registered: true,
	},
	{
		name:       "model/vnd.vtu",
		registered: true,
		extensions: []string{
			"vtu",
		},
	},
	{
		name:       "model/vrml",
		registered: false,
		extensions: []string{
			"wrl", "vrml", "wrz",
		},
	},
	{
		name:       "model/x-pov",
		registered: false,
		extensions: []string{
			"pov",
		},
	},
	{
		name:       "model/x3d+binary",
		registered: false,
		extensions: []string{
			"x3db", "x3dbz",
		},
	},
	{
		name:       "model/x3d+fastinfoset",
		registered: true,
	},
	{
		name:       "model/x3d+vrml",
		registered: false,
		extensions: []string{
			"x3dv", "x3dvz",
		},
	},
	{
		name:       "model/x3d+xml",
		format:     "text/xml",
		registered: true,
		extensions: []string{
			"x3d", "x3dz",
		},
	},
	{
		name:       "model/x3d-vrml",
		registered: true,
	},
	{
		name:       "multipart/alternative",
		registered: false,
	},
	{
		name:       "multipart/appledouble",
		registered: true,
	},
	{
		name:       "multipart/byteranges",
		registered: true,
	},
	{
		name:       "multipart/digest",
		registered: false,
	},
	{
		name:       "multipart/encrypted",
		registered: true,
	},
	{
		name:       "multipart/example",
		registered: true,
	},
	{
		name:       "multipart/form-data",
		registered: true,
	},
	{
		name:       "multipart/header-set",
		registered: true,
	},
	{
		name:       "multipart/mixed",
		registered: false,
	},
	{
		name:       "multipart/multilingual",
		registered: true,
	},
	{
		name:       "multipart/parallel",
		registered: false,
	},
	{
		name:       "multipart/related",
		registered: true,
	},
	{
		name:       "multipart/report",
		registered: true,
	},
	{
		name:       "multipart/signed",
		registered: true,
	},
	{
		name:       "multipart/vnd.bint.med-plus",
		registered: true,
	},
	{
		name:       "multipart/voice-message",
		registered: true,
	},
	{
		name:       "multipart/x-gzip",
		registered: false,
		extensions: []string{
			"gzip",
		},
	},
	{
		name:       "multipart/x-mixed-replace",
		registered: true,
	},
	{
		name:       "multipart/x-ustar",
		registered: false,
		extensions: []string{
			"ustar",
		},
	},
	{
		name:       "multipart/x-zip",
		registered: false,
		extensions: []string{
			"zip",
		},
	},
	{
		name:       "music/crescendo",
		registered: false,
		extensions: []string{
			"mid", "midi",
		},
	},
	{
		name:       "music/x-karaoke",
		registered: false,
		extensions: []string{
			"kar",
		},
	},
	{
		name:       "music/x-midi",
		registered: false,
		extensions: []string{
			"mid", "midi",
		},
	},
	{
		name:       "paleovu/x-pv",
		registered: false,
		extensions: []string{
			"pvu",
		},
	},
	{
		name:       "text/1d-interleaved-parityfec",
		registered: true,
	},
	{
		name:       "text/RED",
		registered: true,
	},
	{
		name:       "text/SGML",
		registered: true,
	},
	{
		name:       "text/asp",
		registered: false,
		extensions: []string{
			"asp",
		},
	},
	{
		name:       "text/cache-manifest",
		registered: true,
		extensions: []string{
			"appcache", "manifest",
		},
	},
	{
		name:       "text/calendar",
		registered: true,
		extensions: []string{
			"ics", "ifb", "icz",
		},
	},
	{
		name:       "text/comma-separated-values",
		registered: false,
		extensions: []string{
			"csv",
		},
	},
	{
		name:       "text/cql",
		registered: true,
	},
	{
		name:       "text/cql-expression",
		registered: true,
	},
	{
		name:       "text/cql-identifier",
		registered: true,
	},
	{
		name:       "text/css",
		registered: true,
		extensions: []string{
			"css",
		},
	},
	{
		name:       "text/csv",
		registered: true,
		extensions: []string{
			"csv",
		},
	},
	{
		name:       "text/csv-schema",
		registered: true,
	},
	{
		name:       "text/directory",
		registered: true,
	},
	{
		name:       "text/dns",
		registered: true,
	},
	{
		name:       "text/ecmascript",
		registered: true,
		extensions: []string{
			"js",
		},
	},
	{
		name:       "text/encaprtp",
		registered: true,
	},
	{
		name:       "text/english",
		registered: false,
	},
	{
		name:       "text/enriched",
		registered: false,
	},
	{
		name:       "text/event-stream",
		registered: false,
		extensions: []string{
			"event-stream",
		},
	},
	{
		name:       "text/example",
		registered: true,
	},
	{
		name:       "text/fhirpath",
		registered: true,
	},
	{
		name:       "text/flexfec",
		registered: true,
	},
	{
		name:       "text/fwdred",
		registered: true,
	},
	{
		name:       "text/gff3",
		registered: true,
	},
	{
		name:       "text/grammar-ref-list",
		registered: true,
	},
	{
		name:       "text/h323",
		registered: false,
		extensions: []string{
			"323",
		},
	},
	{
		name:       "text/hl7v2",
		registered: true,
	},
	{
		name:       "text/html",
		registered: true,
		extensions: []string{
			"html", "acgi", "htm", "htmls", "htx", "shtml", "stm",
		},
	},
	{
		name:       "text/iuls",
		registered: false,
		extensions: []string{
			"uls",
		},
	},
	{
		name:       "text/javascript",
		registered: true,
		extensions: []string{
			"js",
		},
	},
	{
		name:       "text/jcr-cnd",
		registered: true,
	},
	{
		name:       "text/markdown",
		registered: true,
	},
	{
		name:       "text/mathml",
		registered: false,
		extensions: []string{
			"mml",
		},
	},
	{
		name:       "text/mcf",
		registered: false,
		extensions: []string{
			"mcf",
		},
	},
	{
		name:       "text/mizar",
		registered: true,
	},
	{
		name:       "text/n3",
		registered: true,
		extensions: []string{
			"n3",
		},
	},
	{
		name:       "text/parameters",
		registered: true,
	},
	{
		name:       "text/parityfec",
		registered: true,
	},
	{
		name:       "text/pascal",
		registered: false,
		extensions: []string{
			"pas",
		},
	},
	{
		name:       "text/plain",
		registered: false,
		extensions: []string{
			"txt", "text", "conf", "def", "list", "log", "c", "c++", "cc", "com", "cxx", "f", "f90", "for", "g", "h", "hh", "idc", "jav", "java", "lst", "m", "mar", "pl", "sdml", "bas", "in", "asc", "diff", "pot", "el", "ksh",
		},
	},
	{
		name:       "text/plain-bas",
		registered: false,
		extensions: []string{
			"par",
		},
	},
	{
		name:       "text/provenance-notation",
		registered: true,
	},
	{
		name:       "text/prs.fallenstein.rst",
		registered: true,
	},
	{
		name:       "text/prs.lines.tag",
		registered: true,
		extensions: []string{
			"dsc",
		},
	},
	{
		name:       "text/prs.prop.logic",
		registered: true,
	},
	{
		name:       "text/raptorfec",
		registered: true,
	},
	{
		name:       "text/rfc822-headers",
		registered: true,
	},
	{
		name:       "text/richtext",
		registered: false,
		extensions: []string{
			"rtx", "rt", "rtf",
		},
	},
	{
		name:       "text/rtf",
		registered: true,
		extensions: []string{
			"rtf",
		},
	},
	{
		name:       "text/rtp-enc-aescm128",
		registered: true,
	},
	{
		name:       "text/rtploopback",
		registered: true,
	},
	{
		name:       "text/rtx",
		registered: true,
	},
	{
		name:       "text/scriplet",
		registered: false,
		extensions: []string{
			"wsc",
		},
	},
	{
		name:       "text/scriptlet",
		registered: false,
		extensions: []string{
			"sct", "wsc",
		},
	},
	{
		name:       "text/shaclc",
		registered: true,
	},
	{
		name:       "text/shex",
		registered: true,
	},
	{
		name:       "text/spdx",
		registered: true,
	},
	{
		name:       "text/strings",
		registered: true,
	},
	{
		name:       "text/t140",
		registered: true,
	},
	{
		name:       "text/tab-separated-values",
		registered: true,
		extensions: []string{
			"tsv",
		},
	},
	{
		name:       "text/texmacs",
		registered: false,
		extensions: []string{
			"tm", "ts",
		},
	},
	{
		name:       "text/troff",
		registered: true,
		extensions: []string{
			"t", "tr", "roff", "man", "me", "ms",
		},
	},
	{
		name:       "text/turtle",
		registered: true,
		extensions: []string{
			"ttl",
		},
	},
	{
		name:       "text/ulpfec",
		registered: true,
	},
	{
		name:       "text/uri-list",
		registered: true,
		extensions: []string{
			"uri", "uris", "uni", "unis", "urls",
		},
	},
	{
		name:       "text/vcard",
		registered: true,
		extensions: []string{
			"vcard",
		},
	},
	{
		name:       "text/vnd.IPTC.NITF",
		registered: true,
	},
	{
		name:       "text/vnd.IPTC.NewsML",
		registered: true,
	},
	{
		name:       "text/vnd.a",
		registered: true,
	},
	{
		name:       "text/vnd.abc",
		registered: true,
		extensions: []string{
			"abc",
		},
	},
	{
		name:       "text/vnd.ascii-art",
		registered: true,
	},
	{
		name:       "text/vnd.curl",
		registered: true,
		extensions: []string{
			"curl",
		},
	},
	{
		name:       "text/vnd.curl.dcurl",
		registered: false,
		extensions: []string{
			"dcurl",
		},
	},
	{
		name:       "text/vnd.curl.mcurl",
		registered: false,
		extensions: []string{
			"mcurl",
		},
	},
	{
		name:       "text/vnd.curl.scurl",
		registered: false,
		extensions: []string{
			"scurl",
		},
	},
	{
		name:       "text/vnd.debian.copyright",
		registered: true,
	},
	{
		name:       "text/vnd.dmclientscript",
		registered: false,
	},
	{
		name:       "text/vnd.dvb.subtitle",
		registered: true,
		extensions: []string{
			"sub",
		},
	},
	{
		name:       "text/vnd.esmertec.theme-descriptor",
		registered: true,
	},
	{
		name:       "text/vnd.exchangeable",
		registered: true,
	},
	{
		name:       "text/vnd.familysearch.gedcom",
		registered: true,
	},
	{
		name:       "text/vnd.ficlab.flt",
		registered: true,
	},
	{
		name:       "text/vnd.flatland.3dml",
		registered: false,
	},
	{
		name:       "text/vnd.fly",
		registered: true,
		extensions: []string{
			"fly",
		},
	},
	{
		name:       "text/vnd.fmi.flexstor",
		registered: true,
		extensions: []string{
			"flx",
		},
	},
	{
		name:       "text/vnd.gml",
		registered: true,
	},
	{
		name:       "text/vnd.graphviz",
		registered: true,
		extensions: []string{
			"gv",
		},
	},
	{
		name:       "text/vnd.hans",
		registered: true,
	},
	{
		name:       "text/vnd.hgl",
		registered: true,
	},
	{
		name:       "text/vnd.in3d.3dml",
		registered: true,
		extensions: []string{
			"3dml",
		},
	},
	{
		name:       "text/vnd.in3d.spot",
		registered: true,
		extensions: []string{
			"spot",
		},
	},
	{
		name:       "text/vnd.latex-z",
		registered: true,
	},
	{
		name:       "text/vnd.motorola.reflex",
		registered: true,
	},
	{
		name:       "text/vnd.ms-mediapackage",
		registered: true,
	},
	{
		name:       "text/vnd.net2phone.commcenter.command",
		registered: true,
	},
	{
		name:       "text/vnd.radisys.msml-basic-layout",
		registered: true,
	},
	{
		name:       "text/vnd.rn-realtext",
		registered: false,
		extensions: []string{
			"rt",
		},
	},
	{
		name:       "text/vnd.senx.warpscript",
		registered: true,
	},
	{
		name:       "text/vnd.si.uricatalogue",
		registered: true,
	},
	{
		name:       "text/vnd.sosi",
		registered: true,
	},
	{
		name:       "text/vnd.sun.j2me.app-descriptor",
		registered: true,
		extensions: []string{
			"jad",
		},
	},
	{
		name:       "text/vnd.trolltech.linguist",
		registered: true,
	},
	{
		name:       "text/vnd.wap.si",
		registered: true,
		extensions: []string{
			"si",
		},
	},
	{
		name:       "text/vnd.wap.sl",
		registered: true,
		extensions: []string{
			"sl",
		},
	},
	{
		name:       "text/vnd.wap.wml",
		registered: true,
		extensions: []string{
			"wml",
		},
	},
	{
		name:       "text/vnd.wap.wmlscript",
		registered: true,
		extensions: []string{
			"wmls",
		},
	},
	{
		name:       "text/vtt",
		registered: true,
		extensions: []string{
			"vtt",
		},
	},
	{
		name:       "text/webviewhtml",
		registered: false,
		extensions: []string{
			"htt",
		},
	},
	{
		name:       "text/x-asm",
		registered: false,
		extensions: []string{
			"s", "asm",
		},
	},
	{
		name:       "text/x-audiosoft-intra",
		registered: false,
		extensions: []string{
			"aip",
		},
	},
	{
		name:       "text/x-c",
		registered: false,
		extensions: []string{
			"c", "cc", "cxx", "cpp", "h", "hh", "dic",
		},
	},
	{
		name:       "text/x-c++hdr",
		format:     "text/plain",
		registered: false,
		extensions: []string{
			"h++", "hpp", "hxx", "hh",
		},
	},
	{
		name:       "text/x-c++src",
		format:     "text/plain",
		registered: false,
		extensions: []string{
			"c++", "cpp", "cxx", "cc",
		},
	},
	{
		name:       "text/x-chdr",
		registered: false,
		extensions: []string{
			"h",
		},
	},
	{
		name:       "text/x-component",
		registered: false,
		extensions: []string{
			"htc",
		},
	},
	{
		name:       "text/x-crontab",
		registered: false,
	},
	{
		name:       "text/x-csh",
		registered: false,
		extensions: []string{
			"csh",
		},
	},
	{
		name:       "text/x-csrc",
		registered: false,
		extensions: []string{
			"c",
		},
	},
	{
		name:       "text/x-fortran",
		registered: false,
		extensions: []string{
			"f", "for", "f77", "f90",
		},
	},
	{
		name:       "text/x-h",
		registered: false,
		extensions: []string{
			"h", "hh",
		},
	},
	{
		name:       "text/x-java",
		registered: false,
		extensions: []string{
			"java",
		},
	},
	{
		name:       "text/x-java-source",
		registered: false,
		extensions: []string{
			"java", "jav",
		},
	},
	{
		name:       "text/x-la-asf",
		registered: false,
		extensions: []string{
			"lsx",
		},
	},
	{
		name:       "text/x-lua",
		registered: false,
		extensions: []string{
			"lua",
		},
	},
	{
		name:       "text/x-m",
		registered: false,
		extensions: []string{
			"m",
		},
	},
	{
		name:       "text/x-makefile",
		registered: false,
	},
	{
		name:       "text/x-markdown",
		registered: false,
		extensions: []string{
			"markdown", "md", "mkd",
		},
	},
	{
		name:       "text/x-moc",
		registered: false,
		extensions: []string{
			"moc",
		},
	},
	{
		name:       "text/x-nfo",
		registered: false,
		extensions: []string{
			"nfo",
		},
	},
	{
		name:       "text/x-opml",
		registered: false,
		extensions: []string{
			"opml",
		},
	},
	{
		name:       "text/x-pascal",
		registered: false,
		extensions: []string{
			"p", "pas",
		},
	},
	{
		name:       "text/x-pcs-gcd",
		registered: false,
		extensions: []string{
			"gcd",
		},
	},
	{
		name:       "text/x-perl",
		registered: false,
		extensions: []string{
			"pl", "pm",
		},
	},
	{
		name:       "text/x-python",
		registered: false,
		extensions: []string{
			"py",
		},
	},
	{
		name:       "text/x-script",
		registered: false,
		extensions: []string{
			"hlb",
		},
	},
	{
		name:       "text/x-script.csh",
		registered: false,
		extensions: []string{
			"csh",
		},
	},
	{
		name:       "text/x-script.elisp",
		registered: false,
		extensions: []string{
			"el",
		},
	},
	{
		name:       "text/x-script.guile",
		registered: false,
		extensions: []string{
			"scm",
		},
	},
	{
		name:       "text/x-script.ksh",
		registered: false,
		extensions: []string{
			"ksh",
		},
	},
	{
		name:       "text/x-script.lisp",
		registered: false,
		extensions: []string{
			"lsp",
		},
	},
	{
		name:       "text/x-script.perl",
		registered: false,
		extensions: []string{
			"pl",
		},
	},
	{
		name:       "text/x-script.perl-module",
		registered: false,
		extensions: []string{
			"pm",
		},
	},
	{
		name:       "text/x-script.phyton",
		registered: false,
		extensions: []string{
			"py",
		},
	},
	{
		name:       "text/x-script.rexx",
		registered: false,
		extensions: []string{
			"rexx",
		},
	},
	{
		name:       "text/x-script.scheme",
		registered: false,
		extensions: []string{
			"scm",
		},
	},
	{
		name:       "text/x-script.sh",
		registered: false,
		extensions: []string{
			"sh",
		},
	},
	{
		name:       "text/x-script.tcl",
		registered: false,
		extensions: []string{
			"tcl",
		},
	},
	{
		name:       "text/x-script.tcsh",
		registered: false,
		extensions: []string{
			"tcsh",
		},
	},
	{
		name:       "text/x-script.zsh",
		registered: false,
		extensions: []string{
			"zsh",
		},
	},
	{
		name:       "text/x-server-parsed-html",
		registered: false,
		extensions: []string{
			"shtml", "ssi",
		},
	},
	{
		name:       "text/x-setext",
		registered: false,
		extensions: []string{
			"etx",
		},
	},
	{
		name:       "text/x-sfv",
		registered: false,
		extensions: []string{
			"sfv",
		},
	},
	{
		name:       "text/x-sgml",
		registered: false,
		extensions: []string{
			"sgm", "sgml",
		},
	},
	{
		name:       "text/x-sh",
		registered: false,
		extensions: []string{
			"sh",
		},
	},
	{
		name:       "text/x-speech",
		registered: false,
		extensions: []string{
			"spc", "talk",
		},
	},
	{
		name:       "text/x-tcl",
		registered: false,
		extensions: []string{
			"tcl", "tk",
		},
	},
	{
		name:       "text/x-tex",
		registered: false,
		extensions: []string{
			"tex", "ltx", "sty", "cls",
		},
	},
	{
		name:       "text/x-uil",
		registered: false,
		extensions: []string{
			"uil",
		},
	},
	{
		name:       "text/x-uuencode",
		registered: false,
		extensions: []string{
			"uu", "uue",
		},
	},
	{
		name:       "text/x-vcalendar",
		registered: false,
		extensions: []string{
			"vcs",
		},
	},
	{
		name:       "text/x-vcard",
		registered: false,
		extensions: []string{
			"vcf",
		},
	},
	{
		name:       "text/x-yaml",
		registered: false,
		extensions: []string{
			"yaml", "yml",
		},
	},
	{
		name:       "text/xml",
		registered: true,
		extensions: []string{
			"xml",
		},
	},
	{
		name:       "text/xml-external-parsed-entity",
		registered: true,
	},
	{
		name:       "unknown/unknown",
		registered: false,
	},
	{
		name:       "video/1d-interleaved-parityfec",
		registered: true,
	},
	{
		name:       "video/3gpp",
		registered: true,
		extensions: []string{
			"3gp",
		},
	},
	{
		name:       "video/3gpp-tt",
		registered: true,
	},
	{
		name:       "video/3gpp2",
		registered: true,
		extensions: []string{
			"3g2",
		},
	},
	{
		name:       "video/AV1",
		registered: true,
	},
	{
		name:       "video/CelB",
		registered: true,
	},
	{
		name:       "video/DV",
		registered: true,
	},
	{
		name:       "video/FFV1",
		registered: true,
	},
	{
		name:       "video/H261",
		registered: true,
	},
	{
		name:       "video/H263",
		registered: true,
	},
	{
		name:       "video/H263-2000",
		registered: true,
	},
	{
		name:       "video/H265",
		registered: true,
	},
	{
		name:       "video/H266",
		registered: true,
	},
	{
		name:       "video/JPEG",
		registered: true,
	},
	{
		name:       "video/MP1S",
		registered: true,
	},
	{
		name:       "video/MP2P",
		registered: true,
	},
	{
		name:       "video/MP4V-ES",
		registered: true,
	},
	{
		name:       "video/SMPTE292M",
		registered: true,
	},
	{
		name:       "video/VP8",
		registered: true,
	},
	{
		name:       "video/VP9",
		registered: true,
	},
	{
		name:       "video/animaflex",
		registered: false,
		extensions: []string{
			"afl",
		},
	},
	{
		name:       "video/avi",
		registered: false,
		extensions: []string{
			"avi",
		},
	},
	{
		name:       "video/avs-video",
		registered: false,
		extensions: []string{
			"avs",
		},
	},
	{
		name:       "video/bmpeg",
		registered: false,
	},
	{
		name:       "video/bt656",
		registered: false,
	},
	{
		name:       "video/dl",
		registered: false,
		extensions: []string{
			"dl",
		},
	},
	{
		name:       "video/encaprtp",
		registered: true,
	},
	{
		name:       "video/example",
		registered: true,
	},
	{
		name:       "video/flc",
		registered: false,
		extensions: []string{
			"flc", "fli",
		},
	},
	{
		name:       "video/flexfec",
		registered: true,
	},
	{
		name:       "video/fli",
		registered: false,
		extensions: []string{
			"flc", "fli",
		},
	},
	{
		name:       "video/gl",
		registered: false,
		extensions: []string{
			"gl",
		},
	},
	{
		name:       "video/h263-1998",
		registered: false,
	},
	{
		name:       "video/h264",
		registered: false,
		extensions: []string{
			"h264",
		},
	},
	{
		name:       "video/h264-rcdo",
		registered: false,
	},
	{
		name:       "video/h264-svc",
		registered: false,
	},
	{
		name:       "video/iso.segment",
		registered: true,
	},
	{
		name:       "video/jpeg2000",
		registered: true,
	},
	{
		name:       "video/jpm",
		registered: false,
		extensions: []string{
			"jpm", "jpgm",
		},
	},
	{
		name:       "video/jxsv",
		registered: true,
	},
	{
		name:       "video/mj2",
		registered: true,
		extensions: []string{
			"mj2", "mjp2",
		},
	},
	{
		name:       "video/mp2t",
		registered: false,
	},
	{
		name:       "video/mp4",
		registered: true,
		extensions: []string{
			"mp4", "mp4v", "mpg4",
		},
	},
	{
		name:       "video/mpeg",
		registered: false,
		extensions: []string{
			"mpeg", "mpg", "mpe", "m1v", "m2v", "mp2", "mp3", "mpa", "mpv2",
		},
	},
	{
		name:       "video/mpeg4-generic",
		registered: true,
	},
	{
		name:       "video/mpv",
		registered: false,
	},
	{
		name:       "video/msvideo",
		registered: false,
		extensions: []string{
			"avi",
		},
	},
	{
		name:       "video/nv",
		registered: true,
	},
	{
		name:       "video/ogg",
		registered: true,
		extensions: []string{
			"ogv",
		},
	},
	{
		name:       "video/parityfec",
		registered: true,
	},
	{
		name:       "video/pointer",
		registered: true,
	},
	{
		name:       "video/quicktime",
		registered: true,
		extensions: []string{
			"qt", "moov", "mov",
		},
	},
	{
		name:       "video/raptorfec",
		registered: true,
	},
	{
		name:       "video/raw",
		registered: true,
	},
	{
		name:       "video/rtp-enc-aescm128",
		registered: true,
	},
	{
		name:       "video/rtploopback",
		registered: true,
	},
	{
		name:       "video/rtx",
		registered: true,
	},
	{
		name:       "video/scip",
		registered: true,
	},
	{
		name:       "video/smpte291",
		registered: true,
	},
	{
		name:       "video/ulpfec",
		registered: true,
	},
	{
		name:       "video/vc1",
		registered: true,
	},
	{
		name:       "video/vc2",
		registered: true,
	},
	{
		name:       "video/vdo",
		registered: false,
		extensions: []string{
			"vdo",
		},
	},
	{
		name:       "video/vivo",
		registered: false,
		extensions: []string{
			"viv", "vivo",
		},
	},
	{
		name:       "video/vnd.CCTV",
		registered: true,
	},
	{
		name:       "video/vnd.dece.hd",
		registered: true,
		extensions: []string{
			"uvh", "uvvh",
		},
	},
	{
		name:       "video/vnd.dece.mobile",
		registered: true,
		extensions: []string{
			"uvm", "uvvm",
		},
	},
	{
		name:       "video/vnd.dece.mp4",
		registered: true,
	},
	{
		name:       "video/vnd.dece.pd",
		registered: true,
		extensions: []string{
			"uvp", "uvvp",
		},
	},
	{
		name:       "video/vnd.dece.sd",
		registered: true,
		extensions: []string{
			"uvs", "uvvs",
		},
	},
	{
		name:       "video/vnd.dece.video",
		registered: true,
		extensions: []string{
			"uvv", "uvvv",
		},
	},
	{
		name:       "video/vnd.directv.mpeg",
		registered: true,
	},
	{
		name:       "video/vnd.directv.mpeg-tts",
		registered: true,
	},
	{
		name:       "video/vnd.dlna.mpeg-tts",
		registered: true,
	},
	{
		name:       "video/vnd.dvb.file",
		registered: true,
		extensions: []string{
			"dvb",
		},
	},
	{
		name:       "video/vnd.fvt",
		registered: true,
		extensions: []string{
			"fvt",
		},
	},
	{
		name:       "video/vnd.hns.video",
		registered: true,
	},
	{
		name:       "video/vnd.iptvforum.1dparityfec-1010",
		registered: true,
	},
	{
		name:       "video/vnd.iptvforum.1dparityfec-2005",
		registered: true,
	},
	{
		name:       "video/vnd.iptvforum.2dparityfec-1010",
		registered: true,
	},
	{
		name:       "video/vnd.iptvforum.2dparityfec-2005",
		registered: true,
	},
	{
		name:       "video/vnd.iptvforum.ttsavc",
		registered: true,
	},
	{
		name:       "video/vnd.iptvforum.ttsmpeg2",
		registered: true,
	},
	{
		name:       "video/vnd.motorola.video",
		registered: true,
	},
	{
		name:       "video/vnd.motorola.videop",
		registered: true,
	},
	{
		name:       "video/vnd.mpegurl",
		registered: true,
		extensions: []string{
			"mxu", "m4u",
		},
	},
	{
		name:       "video/vnd.ms-playready.media.pyv",
		registered: true,
		extensions: []string{
			"pyv",
		},
	},
	{
		name:       "video/vnd.mts",
		registered: false,
	},
	{
		name:       "video/vnd.nokia.interleaved-multimedia",
		registered: true,
	},
	{
		name:       "video/vnd.nokia.mp4vr",
		registered: true,
	},
	{
		name:       "video/vnd.nokia.videovoip",
		registered: true,
	},
	{
		name:       "video/vnd.objectvideo",
		registered: true,
	},
	{
		name:       "video/vnd.radgamettools.bink",
		registered: true,
	},
	{
		name:       "video/vnd.radgamettools.smacker",
		registered: true,
	},
	{
		name:       "video/vnd.rn-realvideo",
		registered: false,
		extensions: []string{
			"rv",
		},
	},
	{
		name:       "video/vnd.sealed.mpeg1",
		registered: true,
	},
	{
		name:       "video/vnd.sealed.mpeg4",
		registered: true,
	},
	{
		name:       "video/vnd.sealed.swf",
		registered: true,
	},
	{
		name:       "video/vnd.sealedmedia.softseal.mov",
		registered: true,
	},
	{
		name:       "video/vnd.uvvu.mp4",
		registered: true,
		extensions: []string{
			"uvu", "uvvu",
		},
	},
	{
		name:       "video/vnd.vivo",
		registered: true,
		extensions: []string{
			"viv", "vivo",
		},
	},
	{
		name:       "video/vnd.youtube.yt",
		registered: true,
	},
	{
		name:       "video/vosaic",
		registered: false,
		extensions: []string{
			"vos",
		},
	},
	{
		name:       "video/webm",
		registered: false,
		extensions: []string{
			"webm",
		},
	},
	{
		name:       "video/x-amt-demorun",
		registered: false,
		extensions: []string{
			"xdr",
		},
	},
	{
		name:       "video/x-amt-showrun",
		registered: false,
		extensions: []string{
			"xsr",
		},
	},
	{
		name:       "video/x-atomic3d-feature",
		registered: false,
		extensions: []string{
			"fmf",
		},
	},
	{
		name:       "video/x-dl",
		registered: false,
		extensions: []string{
			"dl",
		},
	},
	{
		name:       "video/x-dv",
		registered: false,
		extensions: []string{
			"dif", "dv",
		},
	},
	{
		name:       "video/x-f4v",
		registered: false,
		extensions: []string{
			"f4v",
		},
	},
	{
		name:       "video/x-fli",
		registered: false,
		extensions: []string{
			"fli",
		},
	},
	{
		name:       "video/x-flv",
		registered: false,
		extensions: []string{
			"flv",
		},
	},
	{
		name:       "video/x-gl",
		registered: false,
		extensions: []string{
			"gl",
		},
	},
	{
		name:       "video/x-isvideo",
		registered: false,
		extensions: []string{
			"isu",
		},
	},
	{
		name:       "video/x-la-asf",
		registered: false,
		extensions: []string{
			"lsf", "lsx",
		},
	},
	{
		name:       "video/x-m4v",
		registered: false,
		extensions: []string{
			"m4v",
		},
	},
	{
		name:       "video/x-matroska",
		registered: false,
		extensions: []string{
			"mkv", "mk3d", "mks",
		},
	},
	{
		name:       "video/x-mng",
		registered: false,
		extensions: []string{
			"mng",
		},
	},
	{
		name:       "video/x-motion-jpeg",
		registered: false,
		extensions: []string{
			"mjpg",
		},
	},
	{
		name:       "video/x-mpeg",
		registered: false,
		extensions: []string{
			"mp2", "mp3",
		},
	},
	{
		name:       "video/x-mpeq2a",
		registered: false,
		extensions: []string{
			"mp2",
		},
	},
	{
		name:       "video/x-ms-asf",
		registered: false,
		extensions: []string{
			"asf", "asx", "asr",
		},
	},
	{
		name:       "video/x-ms-asf-plugin",
		registered: false,
		extensions: []string{
			"asx",
		},
	},
	{
		name:       "video/x-ms-vob",
		registered: false,
		extensions: []string{
			"vob",
		},
	},
	{
		name:       "video/x-ms-wm",
		registered: false,
		extensions: []string{
			"wm",
		},
	},
	{
		name:       "video/x-ms-wmv",
		registered: false,
		extensions: []string{
			"wmv",
		},
	},
	{
		name:       "video/x-ms-wmx",
		registered: false,
		extensions: []string{
			"wmx",
		},
	},
	{
		name:       "video/x-ms-wvx",
		registered: false,
		extensions: []string{
			"wvx",
		},
	},
	{
		name:       "video/x-msvideo",
		registered: false,
		extensions: []string{
			"avi",
		},
	},
	{
		name:       "video/x-qtc",
		registered: false,
		extensions: []string{
			"qtc",
		},
	},
	{
		name:       "video/x-scm",
		registered: false,
		extensions: []string{
			"scm",
		},
	},
	{
		name:       "video/x-sgi-movie",
		registered: false,
		extensions: []string{
			"movie", "mv",
		},
	},
	{
		name:       "video/x-smv",
		registered: false,
		extensions: []string{
			"smv",
		},
	},
	{
		name:       "windows/metafile",
		registered: false,
		extensions: []string{
			"wmf",
		},
	},
	{
		name:       "world/i-vrml",
		registered: false,
		extensions: []string{
			"ivr",
		},
	},
	{
		name:       "world/x-3dmf",
		registered: false,
		extensions: []string{
			"3dm", "3dmf", "qd3", "qd3d",
		},
	},
	{
		name:       "world/x-svr",
		registered: false,
		extensions: []string{
			"svr",
		},
	},
	{
		name:       "world/x-vrml",
		registered: false,
		extensions: []string{
			"vrml", "wrl", "wrz", "flr", "xaf", "xof", "vrm",
		},
	},
	{
		name:       "world/x-vrt",
		registered: false,
		extensions: []string{
			"vrt",
		},
	},
	{
		name:       "www/mime",
		registered: false,
		extensions: []string{
			"mime",
		},
	},
	{
		name:       "x-conference/x-cooltalk",
		registered: false,
		extensions: []string{
			"ice",
		},
	},
	{
		name:       "xgl/drawing",
		registered: false,
		extensions: []string{
			"xgz",
		},
	},
	{
		name:       "xgl/movie",
		registered: false,
		extensions: []string{
			"xmz",
		},
	},
}

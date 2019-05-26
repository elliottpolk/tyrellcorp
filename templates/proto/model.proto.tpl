// Code generated by tyrellcorp. DO NOT EDIT.
// source: model.proto.tpl

syntax = "proto3";
package {{ .Package | ToLower | Trim }};

import "record.proto";

option go_package = "{{ .Repository | ToLower | Trim }}/{{ .Package | ToLower | Trim }}";
option java_multiple_files = true;
option java_outer_classname = "{{ .Name }}Proto";
option java_package = "com.tyrellcorp.{{ .Package | ToLower | Trim }}";

message {{ .Name }} {
	// standard record values
	{{ .Package | ToLower | Trim }}.RecordInfo record_info = 1;

    // unique identifier
    string id = 2;

	{{ range .Fields }}
	// {{ .Description }}
	{{ if .IsList }}repeated {{ end }}{{ .Type }} {{ .Name }} = {{ .Sequence }};
	{{ end }}
}

// Code generated by tyrellcorp. DO NOT EDIT.
// source: service.proto.tpl

syntax = "proto3";
package {{ .Package | ToLower | Trim }};

import "{{ .Name | ToLower | Trim }}.proto";
import "google/api/annotations.proto";

option go_package = "{{ .Repository | ToLower | Trim }}/{{ .Package | ToLower | Trim }}";
option java_multiple_files = true;
option java_outer_classname = "{{ .Name }}ServiceProto";
option java_package = "com.tyrellcorp.{{ .Package | ToLower | Trim }}";

message Empty {
  // unique identifier of the original incoming request to help troubleshoot
  string request_id = 1;
}

message {{ .Name }}Request {
  // unique identifier to help troubleshoot each request
  string request_id = 1;

  // username of the one making the request
  string username = 2;

  // unique identifier of the {{ .Package | ToLower | Trim }}.{{ .Name }}
  string id = 3;

  // dataset to process
  repeated {{ .Package | ToLower | Trim }}.{{ .Name }} payload = 4;
}

message {{ .Name }}Response {
  // unique identifier of the original incoming request to help troubleshoot
  string request_id = 1;

  repeated {{ .Package | ToLower | Trim }}.{{ .Name }} payload = 2;
}

service {{ .Name }}Service {
  {{ if .Create }}
  // create new {{ .Name }} item(s)
  rpc Create({{ .Name }}Request) returns (Empty) {
    option (google.api.http) = {
      post: "/api/v1/{{ .Name | ToLower | Trim }}s"
      body: "*"
    };
  }
  {{ end }}
  {{ if .Retrieve }}
  // retrieve a list of {{ .Name }} items
  rpc Retrieve({{ .Name }}Request) returns ({{ .Name }}Response) {
    option (google.api.http) = {
      get: "/api/v1/{{ .Name | ToLower | Trim }}s"

      additional_bindings {
        get: "/api/v1/{{ .Name | ToLower | Trim }}s/{id}"
      }
    };
  }
  {{ end }}
  {{ if .Update }}
  // update {{ .Name }} item(s)
  rpc Update({{ .Name }}Request) returns ({{ .Name }}Response) {
    option (google.api.http) = {
      put: "/api/v1/{{ .Name | ToLower | Trim }}s/{id}"
      body: "*"
    };
  }
  {{ end }}
  {{ if .Delete }}
  // delete {{ .Name }} item(s)
  rpc Delete({{ .Name }}Request) returns (Empty) {
    option (google.api.http) = {
      delete: "/api/v1/{{ .Name | ToLower | Trim }}s"
      body: "*"
    };
  }
  {{ end }}
}
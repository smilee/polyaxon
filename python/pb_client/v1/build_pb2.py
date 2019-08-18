#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import sys


def main(argv):
    pass


if __name__ == '__main__':
    main(sys.argv)
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/build.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from protoc_gen_swagger.options import annotations_pb2 as protoc__gen__swagger_dot_options_dot_annotations__pb2
from v1 import base_pb2 as v1_dot_base__pb2
from v1 import code_ref_pb2 as v1_dot_code__ref__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='v1/build.proto',
  package='v1',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0ev1/build.proto\x12\x02v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a,protoc-gen-swagger/options/annotations.proto\x1a\rv1/base.proto\x1a\x11v1/code_ref.proto\"\xe6\x03\n\x05\x42uild\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04uuid\x18\x02 \x01(\t\x12\x13\n\x0bunique_name\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x0c\n\x04tags\x18\x06 \x03(\t\x12\x0f\n\x07\x64\x65leted\x18\x07 \x01(\x08\x12\x0c\n\x04user\x18\x08 \x01(\t\x12\x12\n\ncreated_at\x18\t \x01(\t\x12\x12\n\nupdated_at\x18\n \x01(\t\x12\x12\n\nstarted_at\x18\x0b \x01(\t\x12\x13\n\x0b\x66inished_at\x18\x0c \x01(\t\x12\x0f\n\x07project\x18\r \x01(\t\x12\x12\n\nis_managed\x18\x0e \x01(\t\x12\x0c\n\x04spec\x18\x0f \x01(\t\x12\x0f\n\x07\x62\x61\x63kend\x18\x10 \x01(\t\x12\x11\n\tframework\x18\x11 \x01(\t\x12\x13\n\x0blast_status\x18\x12 \x01(\t\x12\x16\n\x0e\x63ode_reference\x18\x13 \x01(\x03\x12\x1b\n\tresources\x18\x14 \x01(\x0b\x32\x08.v1.Dict\x12\x0e\n\x06readme\x18\x15 \x01(\t\x12\x12\n\nbookmarked\x18\x16 \x01(\x08\x12\x18\n\x06params\x18\x17 \x01(\x0b\x32\x08.v1.Dict\x12\x19\n\x07run_env\x18\x18 \x01(\x0b\x32\x08.v1.Dict\x12\x13\n\x0b\x62uild_build\x18\x19 \x01(\t\"p\n\x0b\x42uildStatus\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04uuid\x18\x02 \x01(\t\x12\x12\n\ncreated_at\x18\x03 \x01(\t\x12\x12\n\nupdated_at\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\x12\x0f\n\x07message\x18\x06 \x01(\t\"L\n\x10\x42uildBodyRequest\x12\r\n\x05owner\x18\x01 \x01(\t\x12\x0f\n\x07project\x18\x02 \x01(\t\x12\x18\n\x05\x62uild\x18\x03 \x01(\x0b\x32\t.v1.Build\"_\n\x12ListBuildsResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12\x1a\n\x07results\x18\x02 \x03(\x0b\x32\t.v1.Build\x12\x10\n\x08previous\x18\x03 \x01(\t\x12\x0c\n\x04next\x18\x04 \x01(\t\"l\n\x19ListBuildStatusesResponse\x12\r\n\x05\x63ount\x18\x01 \x01(\x05\x12 \n\x07results\x18\x02 \x03(\x0b\x32\x0f.v1.BuildStatus\x12\x10\n\x08previous\x18\x03 \x01(\t\x12\x0c\n\x04next\x18\x04 \x01(\tb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,protoc__gen__swagger_dot_options_dot_annotations__pb2.DESCRIPTOR,v1_dot_base__pb2.DESCRIPTOR,v1_dot_code__ref__pb2.DESCRIPTOR,])




_BUILD = _descriptor.Descriptor(
  name='Build',
  full_name='v1.Build',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='v1.Build.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uuid', full_name='v1.Build.uuid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unique_name', full_name='v1.Build.unique_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='v1.Build.name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='v1.Build.description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='v1.Build.tags', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='deleted', full_name='v1.Build.deleted', index=6,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user', full_name='v1.Build.user', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='v1.Build.created_at', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.Build.updated_at', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='started_at', full_name='v1.Build.started_at', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='finished_at', full_name='v1.Build.finished_at', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project', full_name='v1.Build.project', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_managed', full_name='v1.Build.is_managed', index=13,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='v1.Build.spec', index=14,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='backend', full_name='v1.Build.backend', index=15,
      number=16, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='framework', full_name='v1.Build.framework', index=16,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='last_status', full_name='v1.Build.last_status', index=17,
      number=18, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='code_reference', full_name='v1.Build.code_reference', index=18,
      number=19, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='v1.Build.resources', index=19,
      number=20, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='readme', full_name='v1.Build.readme', index=20,
      number=21, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bookmarked', full_name='v1.Build.bookmarked', index=21,
      number=22, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='params', full_name='v1.Build.params', index=22,
      number=23, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='run_env', full_name='v1.Build.run_env', index=23,
      number=24, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='build_build', full_name='v1.Build.build_build', index=24,
      number=25, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=162,
  serialized_end=648,
)


_BUILDSTATUS = _descriptor.Descriptor(
  name='BuildStatus',
  full_name='v1.BuildStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='v1.BuildStatus.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uuid', full_name='v1.BuildStatus.uuid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='v1.BuildStatus.created_at', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.BuildStatus.updated_at', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='v1.BuildStatus.status', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='v1.BuildStatus.message', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=650,
  serialized_end=762,
)


_BUILDBODYREQUEST = _descriptor.Descriptor(
  name='BuildBodyRequest',
  full_name='v1.BuildBodyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='owner', full_name='v1.BuildBodyRequest.owner', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project', full_name='v1.BuildBodyRequest.project', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='build', full_name='v1.BuildBodyRequest.build', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=764,
  serialized_end=840,
)


_LISTBUILDSRESPONSE = _descriptor.Descriptor(
  name='ListBuildsResponse',
  full_name='v1.ListBuildsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='v1.ListBuildsResponse.count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='v1.ListBuildsResponse.results', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='previous', full_name='v1.ListBuildsResponse.previous', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next', full_name='v1.ListBuildsResponse.next', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=842,
  serialized_end=937,
)


_LISTBUILDSTATUSESRESPONSE = _descriptor.Descriptor(
  name='ListBuildStatusesResponse',
  full_name='v1.ListBuildStatusesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='count', full_name='v1.ListBuildStatusesResponse.count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='results', full_name='v1.ListBuildStatusesResponse.results', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='previous', full_name='v1.ListBuildStatusesResponse.previous', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next', full_name='v1.ListBuildStatusesResponse.next', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=939,
  serialized_end=1047,
)

_BUILD.fields_by_name['resources'].message_type = v1_dot_base__pb2._DICT
_BUILD.fields_by_name['params'].message_type = v1_dot_base__pb2._DICT
_BUILD.fields_by_name['run_env'].message_type = v1_dot_base__pb2._DICT
_BUILDBODYREQUEST.fields_by_name['build'].message_type = _BUILD
_LISTBUILDSRESPONSE.fields_by_name['results'].message_type = _BUILD
_LISTBUILDSTATUSESRESPONSE.fields_by_name['results'].message_type = _BUILDSTATUS
DESCRIPTOR.message_types_by_name['Build'] = _BUILD
DESCRIPTOR.message_types_by_name['BuildStatus'] = _BUILDSTATUS
DESCRIPTOR.message_types_by_name['BuildBodyRequest'] = _BUILDBODYREQUEST
DESCRIPTOR.message_types_by_name['ListBuildsResponse'] = _LISTBUILDSRESPONSE
DESCRIPTOR.message_types_by_name['ListBuildStatusesResponse'] = _LISTBUILDSTATUSESRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Build = _reflection.GeneratedProtocolMessageType('Build', (_message.Message,), {
  'DESCRIPTOR' : _BUILD,
  '__module__' : 'v1.build_pb2'
  # @@protoc_insertion_point(class_scope:v1.Build)
  })
_sym_db.RegisterMessage(Build)

BuildStatus = _reflection.GeneratedProtocolMessageType('BuildStatus', (_message.Message,), {
  'DESCRIPTOR' : _BUILDSTATUS,
  '__module__' : 'v1.build_pb2'
  # @@protoc_insertion_point(class_scope:v1.BuildStatus)
  })
_sym_db.RegisterMessage(BuildStatus)

BuildBodyRequest = _reflection.GeneratedProtocolMessageType('BuildBodyRequest', (_message.Message,), {
  'DESCRIPTOR' : _BUILDBODYREQUEST,
  '__module__' : 'v1.build_pb2'
  # @@protoc_insertion_point(class_scope:v1.BuildBodyRequest)
  })
_sym_db.RegisterMessage(BuildBodyRequest)

ListBuildsResponse = _reflection.GeneratedProtocolMessageType('ListBuildsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTBUILDSRESPONSE,
  '__module__' : 'v1.build_pb2'
  # @@protoc_insertion_point(class_scope:v1.ListBuildsResponse)
  })
_sym_db.RegisterMessage(ListBuildsResponse)

ListBuildStatusesResponse = _reflection.GeneratedProtocolMessageType('ListBuildStatusesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTBUILDSTATUSESRESPONSE,
  '__module__' : 'v1.build_pb2'
  # @@protoc_insertion_point(class_scope:v1.ListBuildStatusesResponse)
  })
_sym_db.RegisterMessage(ListBuildStatusesResponse)


# @@protoc_insertion_point(module_scope)

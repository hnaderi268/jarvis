from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class TextRequest(_message.Message):
    __slots__ = ["text", "max_length", "num_return_sequences"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    MAX_LENGTH_FIELD_NUMBER: _ClassVar[int]
    NUM_RETURN_SEQUENCES_FIELD_NUMBER: _ClassVar[int]
    text: str
    max_length: int
    num_return_sequences: int
    def __init__(self, text: _Optional[str] = ..., max_length: _Optional[int] = ..., num_return_sequences: _Optional[int] = ...) -> None: ...

class TextResponse(_message.Message):
    __slots__ = ["text"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, text: _Optional[_Iterable[str]] = ...) -> None: ...

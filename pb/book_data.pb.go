// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: proto/book_data.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublishedDate string                 `protobuf:"bytes,4,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	UserId        int32                  `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_proto_book_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *Book) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Book) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BorrowedBook struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BorrowedDate  string                 `protobuf:"bytes,3,opt,name=borrowed_date,json=borrowedDate,proto3" json:"borrowed_date,omitempty"`
	ReturnDate    string                 `protobuf:"bytes,4,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BorrowedBook) Reset() {
	*x = BorrowedBook{}
	mi := &file_proto_book_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowedBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowedBook) ProtoMessage() {}

func (x *BorrowedBook) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowedBook.ProtoReflect.Descriptor instead.
func (*BorrowedBook) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{1}
}

func (x *BorrowedBook) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BorrowedBook) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BorrowedBook) GetBorrowedDate() string {
	if x != nil {
		return x.BorrowedDate
	}
	return ""
}

func (x *BorrowedBook) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

// Message untuk permintaan kosong (placeholder)
type EmptyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_proto_book_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{2}
}

// List of books response
type BooksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Books         []*Book                `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BooksResponse) Reset() {
	*x = BooksResponse{}
	mi := &file_proto_book_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksResponse) ProtoMessage() {}

func (x *BooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksResponse.ProtoReflect.Descriptor instead.
func (*BooksResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{3}
}

func (x *BooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

// Endpoint for getting a book by ID
type GetBookByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookByIdRequest) Reset() {
	*x = GetBookByIdRequest{}
	mi := &file_proto_book_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookByIdRequest) ProtoMessage() {}

func (x *GetBookByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookByIdRequest.ProtoReflect.Descriptor instead.
func (*GetBookByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookByIdRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type GetBookByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Book          *Book                  `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookByIdResponse) Reset() {
	*x = GetBookByIdResponse{}
	mi := &file_proto_book_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookByIdResponse) ProtoMessage() {}

func (x *GetBookByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookByIdResponse.ProtoReflect.Descriptor instead.
func (*GetBookByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookByIdResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// Endpoint for Borrowing a book
// Request for borrowing a book
type BorrowBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowedDate  string                 `protobuf:"bytes,2,opt,name=borrowed_date,json=borrowedDate,proto3" json:"borrowed_date,omitempty"`
	ReturnDate    string                 `protobuf:"bytes,3,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"` //  int32 user_id = 4;
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BorrowBookRequest) Reset() {
	*x = BorrowBookRequest{}
	mi := &file_proto_book_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookRequest) ProtoMessage() {}

func (x *BorrowBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookRequest.ProtoReflect.Descriptor instead.
func (*BorrowBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{6}
}

func (x *BorrowBookRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BorrowBookRequest) GetBorrowedDate() string {
	if x != nil {
		return x.BorrowedDate
	}
	return ""
}

func (x *BorrowBookRequest) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

// Response for borrowing a book
type BorrowBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BorrowBookResponse) Reset() {
	*x = BorrowBookResponse{}
	mi := &file_proto_book_data_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookResponse) ProtoMessage() {}

func (x *BorrowBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookResponse.ProtoReflect.Descriptor instead.
func (*BorrowBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{7}
}

func (x *BorrowBookResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BorrowBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Endpoint for Returning a book
type ReturnBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReturnBookRequest) Reset() {
	*x = ReturnBookRequest{}
	mi := &file_proto_book_data_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReturnBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnBookRequest) ProtoMessage() {}

func (x *ReturnBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnBookRequest.ProtoReflect.Descriptor instead.
func (*ReturnBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{8}
}

func (x *ReturnBookRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

// Endpoint for creating a book
type CreateBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PublishedDate string                 `protobuf:"bytes,3,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	mi := &file_proto_book_data_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{9}
}

func (x *CreateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBookRequest) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *CreateBookRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	mi := &file_proto_book_data_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{10}
}

func (x *CreateBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Endpoint for updating a book
type UpdateBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublishedDate string                 `protobuf:"bytes,4,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	mi := &file_proto_book_data_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateBookRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *UpdateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdateBookRequest) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *UpdateBookRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookResponse) Reset() {
	*x = UpdateBookResponse{}
	mi := &file_proto_book_data_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookResponse) ProtoMessage() {}

func (x *UpdateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Endpoint for deleting a book
type DeleteBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        int32                  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	mi := &file_proto_book_data_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteBookRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type DeleteBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBookResponse) Reset() {
	*x = DeleteBookResponse{}
	mi := &file_proto_book_data_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResponse) ProtoMessage() {}

func (x *DeleteBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_data_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_data_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_book_data_proto protoreflect.FileDescriptor

var file_proto_book_data_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x36, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x72, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22,
	0x80, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc2, 0x03, 0x0a,
	0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_book_data_proto_rawDescOnce sync.Once
	file_proto_book_data_proto_rawDescData = file_proto_book_data_proto_rawDesc
)

func file_proto_book_data_proto_rawDescGZIP() []byte {
	file_proto_book_data_proto_rawDescOnce.Do(func() {
		file_proto_book_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_book_data_proto_rawDescData)
	})
	return file_proto_book_data_proto_rawDescData
}

var file_proto_book_data_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_book_data_proto_goTypes = []any{
	(*Book)(nil),                // 0: book_data.Book
	(*BorrowedBook)(nil),        // 1: book_data.BorrowedBook
	(*EmptyRequest)(nil),        // 2: book_data.EmptyRequest
	(*BooksResponse)(nil),       // 3: book_data.BooksResponse
	(*GetBookByIdRequest)(nil),  // 4: book_data.GetBookByIdRequest
	(*GetBookByIdResponse)(nil), // 5: book_data.GetBookByIdResponse
	(*BorrowBookRequest)(nil),   // 6: book_data.BorrowBookRequest
	(*BorrowBookResponse)(nil),  // 7: book_data.BorrowBookResponse
	(*ReturnBookRequest)(nil),   // 8: book_data.ReturnBookRequest
	(*CreateBookRequest)(nil),   // 9: book_data.CreateBookRequest
	(*CreateBookResponse)(nil),  // 10: book_data.CreateBookResponse
	(*UpdateBookRequest)(nil),   // 11: book_data.UpdateBookRequest
	(*UpdateBookResponse)(nil),  // 12: book_data.UpdateBookResponse
	(*DeleteBookRequest)(nil),   // 13: book_data.DeleteBookRequest
	(*DeleteBookResponse)(nil),  // 14: book_data.DeleteBookResponse
}
var file_proto_book_data_proto_depIdxs = []int32{
	0,  // 0: book_data.BooksResponse.books:type_name -> book_data.Book
	0,  // 1: book_data.GetBookByIdResponse.book:type_name -> book_data.Book
	2,  // 2: book_data.BookService.GetBooks:input_type -> book_data.EmptyRequest
	4,  // 3: book_data.BookService.GetBook:input_type -> book_data.GetBookByIdRequest
	9,  // 4: book_data.BookService.CreateBook:input_type -> book_data.CreateBookRequest
	11, // 5: book_data.BookService.UpdateBook:input_type -> book_data.UpdateBookRequest
	13, // 6: book_data.BookService.DeleteBook:input_type -> book_data.DeleteBookRequest
	6,  // 7: book_data.BookService.BorrowBook:input_type -> book_data.BorrowBookRequest
	3,  // 8: book_data.BookService.GetBooks:output_type -> book_data.BooksResponse
	5,  // 9: book_data.BookService.GetBook:output_type -> book_data.GetBookByIdResponse
	10, // 10: book_data.BookService.CreateBook:output_type -> book_data.CreateBookResponse
	12, // 11: book_data.BookService.UpdateBook:output_type -> book_data.UpdateBookResponse
	14, // 12: book_data.BookService.DeleteBook:output_type -> book_data.DeleteBookResponse
	7,  // 13: book_data.BookService.BorrowBook:output_type -> book_data.BorrowBookResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_book_data_proto_init() }
func file_proto_book_data_proto_init() {
	if File_proto_book_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_book_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_book_data_proto_goTypes,
		DependencyIndexes: file_proto_book_data_proto_depIdxs,
		MessageInfos:      file_proto_book_data_proto_msgTypes,
	}.Build()
	File_proto_book_data_proto = out.File
	file_proto_book_data_proto_rawDesc = nil
	file_proto_book_data_proto_goTypes = nil
	file_proto_book_data_proto_depIdxs = nil
}

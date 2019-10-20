// package: priceMonitor
// file: services/priceMonitorpb/priceMonitor.proto

import * as jspb from "google-protobuf";

export class Product extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  getImages(): string;
  setImages(value: string): void;

  clearHistoryList(): void;
  getHistoryList(): Array<PriceTime>;
  setHistoryList(value: Array<PriceTime>): void;
  addHistory(value?: PriceTime, index?: number): PriceTime;

  getCreatedAt(): string;
  setCreatedAt(value: string): void;

  getName(): string;
  setName(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getPrice(): string;
  setPrice(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Product.AsObject;
  static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Product;
  static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
  export type AsObject = {
    id: string,
    url: string,
    images: string,
    historyList: Array<PriceTime.AsObject>,
    createdAt: string,
    name: string,
    description: string,
    price: string,
  }
}

export class PriceTime extends jspb.Message {
  getPrice(): string;
  setPrice(value: string): void;

  getTime(): string;
  setTime(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PriceTime.AsObject;
  static toObject(includeInstance: boolean, msg: PriceTime): PriceTime.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PriceTime, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PriceTime;
  static deserializeBinaryFromReader(message: PriceTime, reader: jspb.BinaryReader): PriceTime;
}

export namespace PriceTime {
  export type AsObject = {
    price: string,
    time: string,
  }
}

export class AddProductRequest extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddProductRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddProductRequest): AddProductRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddProductRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddProductRequest;
  static deserializeBinaryFromReader(message: AddProductRequest, reader: jspb.BinaryReader): AddProductRequest;
}

export namespace AddProductRequest {
  export type AsObject = {
    url: string,
  }
}

export class AddProductResponse extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddProductResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddProductResponse): AddProductResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddProductResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddProductResponse;
  static deserializeBinaryFromReader(message: AddProductResponse, reader: jspb.BinaryReader): AddProductResponse;
}

export namespace AddProductResponse {
  export type AsObject = {
    product?: Product.AsObject,
  }
}

export class GetProductRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProductRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProductRequest): GetProductRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProductRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProductRequest;
  static deserializeBinaryFromReader(message: GetProductRequest, reader: jspb.BinaryReader): GetProductRequest;
}

export namespace GetProductRequest {
  export type AsObject = {
    id: string,
  }
}

export class ProductResponse extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProductResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ProductResponse): ProductResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProductResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProductResponse;
  static deserializeBinaryFromReader(message: ProductResponse, reader: jspb.BinaryReader): ProductResponse;
}

export namespace ProductResponse {
  export type AsObject = {
    product?: Product.AsObject,
  }
}

export class GetProductsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProductsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProductsRequest): GetProductsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProductsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProductsRequest;
  static deserializeBinaryFromReader(message: GetProductsRequest, reader: jspb.BinaryReader): GetProductsRequest;
}

export namespace GetProductsRequest {
  export type AsObject = {
  }
}

export class ProductsResponse extends jspb.Message {
  clearProductsList(): void;
  getProductsList(): Array<Product>;
  setProductsList(value: Array<Product>): void;
  addProducts(value?: Product, index?: number): Product;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProductsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ProductsResponse): ProductsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProductsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProductsResponse;
  static deserializeBinaryFromReader(message: ProductsResponse, reader: jspb.BinaryReader): ProductsResponse;
}

export namespace ProductsResponse {
  export type AsObject = {
    productsList: Array<Product.AsObject>,
  }
}


// package: priceMonitor
// file: services/priceMonitorpb/priceMonitor.proto

import * as services_priceMonitorpb_priceMonitor_pb from "../../proto/priceMonitorpb/priceMonitor_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PriceMonitorServiceAddProduct = {
  readonly methodName: string;
  readonly service: typeof PriceMonitorService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof services_priceMonitorpb_priceMonitor_pb.AddProductRequest;
  readonly responseType: typeof services_priceMonitorpb_priceMonitor_pb.AddProductResponse;
};

type PriceMonitorServiceGetProduct = {
  readonly methodName: string;
  readonly service: typeof PriceMonitorService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof services_priceMonitorpb_priceMonitor_pb.GetProductRequest;
  readonly responseType: typeof services_priceMonitorpb_priceMonitor_pb.ProductResponse;
};

type PriceMonitorServiceGetProducts = {
  readonly methodName: string;
  readonly service: typeof PriceMonitorService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof services_priceMonitorpb_priceMonitor_pb.GetProductsRequest;
  readonly responseType: typeof services_priceMonitorpb_priceMonitor_pb.ProductsResponse;
};

export class PriceMonitorService {
  static readonly serviceName: string;
  static readonly AddProduct: PriceMonitorServiceAddProduct;
  static readonly GetProduct: PriceMonitorServiceGetProduct;
  static readonly GetProducts: PriceMonitorServiceGetProducts;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class PriceMonitorServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  addProduct(
    requestMessage: services_priceMonitorpb_priceMonitor_pb.AddProductRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: services_priceMonitorpb_priceMonitor_pb.AddProductResponse|null) => void
  ): UnaryResponse;
  addProduct(
    requestMessage: services_priceMonitorpb_priceMonitor_pb.AddProductRequest,
    callback: (error: ServiceError|null, responseMessage: services_priceMonitorpb_priceMonitor_pb.AddProductResponse|null) => void
  ): UnaryResponse;
  getProduct(
    requestMessage: services_priceMonitorpb_priceMonitor_pb.GetProductRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: services_priceMonitorpb_priceMonitor_pb.ProductResponse|null) => void
  ): UnaryResponse;
  getProduct(
    requestMessage: services_priceMonitorpb_priceMonitor_pb.GetProductRequest,
    callback: (error: ServiceError|null, responseMessage: services_priceMonitorpb_priceMonitor_pb.ProductResponse|null) => void
  ): UnaryResponse;
  getProducts(
    requestMessage: services_priceMonitorpb_priceMonitor_pb.GetProductsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: services_priceMonitorpb_priceMonitor_pb.ProductsResponse|null) => void
  ): UnaryResponse;
  getProducts(
    requestMessage: services_priceMonitorpb_priceMonitor_pb.GetProductsRequest,
    callback: (error: ServiceError|null, responseMessage: services_priceMonitorpb_priceMonitor_pb.ProductsResponse|null) => void
  ): UnaryResponse;
}


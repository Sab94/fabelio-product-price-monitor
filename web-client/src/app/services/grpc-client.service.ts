import { Injectable } from '@angular/core';

import {
  AddProductRequest,
  GetProductRequest,
  GetProductsRequest,
  ProductResponse,
  ProductsResponse
} from '../proto/priceMonitorpb/priceMonitor_pb';
import { PriceMonitorServiceClient } from '../proto/priceMonitorpb/priceMonitor_pb_service';
@Injectable({
  providedIn: 'root'
})
export class GRPCClient {
  client: PriceMonitorServiceClient;
  constructor() {
    this.client = new PriceMonitorServiceClient('http://ec2-18-219-84-250.us-east-2.compute.amazonaws.com:8080');
  }

  addProduct(url) {
    return new Promise(
      function(resolve, reject) {
        const req = new AddProductRequest()
        req.setUrl(url);
        this.client.addProduct(req, null, (err, response) => {
          if (err) {
            reject(err);
          } else {
            resolve(response);
          }
        });
      }.bind(this)
    );
  }

  getProducts() {
    return new Promise(
      function(resolve, reject) {
        const req = new GetProductsRequest();
        this.client.getProducts(req, null, (err, response: ProductsResponse) => {
          if (err) {
            reject(err);
          } else {
            resolve(response);
          }
        });
      }.bind(this)
    );
  }

  getProduct(id) {
    return new Promise(
      function(resolve, reject) {
        const req = new GetProductRequest();
        req.setId(id)
        this.client.getProduct(req, null, (err, response: ProductResponse) => {
          if (err) {
            reject(err);
          } else {
            resolve(response);
          }
        });
      }.bind(this)
    );
  }
}

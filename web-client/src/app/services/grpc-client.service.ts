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
    this.client = new PriceMonitorServiceClient('http://localhost:8000');
  }

  addProduct(url) {
    return new Promise(
      function(resolve, reject) {
        const req = new AddProductRequest()
        req.setUrl(url);
        this.client.addProduct(req, null, (err, response) => {
          console.log('a')
          if (err) {
            reject(err);
          } else {
            console.log('b', response)
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
          console.log('a')
          if (err) {
            reject(err);
          } else {
            console.log('b', response)
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
          console.log('a')
          if (err) {
            reject(err);
          } else {
            console.log('b', response)
            resolve(response);
          }
        });
      }.bind(this)
    );
  }
}

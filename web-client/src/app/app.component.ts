import { Component } from '@angular/core';
import { AddProductRequest  } from './proto/priceMonitorpb/priceMonitor_pb';
import { PriceMonitorServiceClient } from './proto/priceMonitorpb/priceMonitor_pb_service';
import {grpc} from '@improbable-eng/grpc-web';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'web-client';
  client: PriceMonitorServiceClient;
  constructor() {
    const req = new AddProductRequest()
    req.setUrl('https://fabelio.com/ip/jon-shelf.html')
    
    this.client = new PriceMonitorServiceClient('http://localhost:8000');
    this.client.addProduct(req, null, (err, response) => {
      if(err) {
        console.log(err)
      } else {
        console.log('ApiService.list.response', response);
      }
    });
  }

}

import { Component, OnInit } from '@angular/core';
import {ProductResponse} from '../../proto/priceMonitorpb/priceMonitor_pb';

import {GRPCClient} from '../../services/grpc-client.service';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.scss']
})
export class DetailComponent implements OnInit {
  product: any
  constructor(private grpcClient: GRPCClient, private route: ActivatedRoute) {
  }

  get(id) {
    this.grpcClient.getProduct(id).then((res: ProductResponse) => {
      this.product = res.toObject().product
      console.log(this.product)
    })
  }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.get(params['id'])
    });
  }

}

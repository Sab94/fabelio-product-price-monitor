import { Component, OnInit } from '@angular/core';
import {ProductResponse} from '../../proto/priceMonitorpb/priceMonitor_pb';
import * as moment from 'moment';

import {GRPCClient} from '../../services/grpc-client.service';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.scss']
})
export class DetailComponent implements OnInit {
  product: any;
  myDataSets = [{
    name: 'Price',
    series: []
  }];

  constructor(private grpcClient: GRPCClient, private route: ActivatedRoute) {
  }

  get(id) {
    this.grpcClient.getProduct(id).then((res: ProductResponse) => {
      this.product = res.toObject().product
      for (let i = 0; i < this.product.historyList.length; i++) {
        this.myDataSets[0].series.push({
          name: this.product.historyList[i].time,
          value: this.product.historyList[i].price,
        })
      }
    })
  }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.get(params['id'])
    });
  }

  dateTickFormatting(val: any): string {
    return moment(val, "YYYY-MM-DD HH:mm:ss.SSSS").format('LT L')
  }

}

import { Component, OnInit } from '@angular/core';
import {GRPCClient} from '../../services/grpc-client.service';
import {Router} from '@angular/router';
import {ProductResponse, ProductsResponse} from '../../proto/priceMonitorpb/priceMonitor_pb';

@Component({
  selector: 'app-all',
  templateUrl: './all.component.html',
  styleUrls: ['./all.component.scss']
})
export class AllComponent implements OnInit {
  products: any;
  constructor(private grpcClient: GRPCClient, private router: Router) {
  }

  get() {
    this.grpcClient.getProducts().then((res: ProductsResponse) => {
      this.products = res.toObject().productsList
      console.log(this.products)
    })
  }

  ngOnInit(): void {
    this.get();
  }

  gotoProduct(id) {
    console.log(id)
    this.router.navigateByUrl('/detail/' + id);
  }
}

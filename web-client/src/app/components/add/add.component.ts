import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';

import {GRPCClient} from '../../services/grpc-client.service';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.scss']
})
export class AddComponent {
  url: string;
  loading: any  = false;
  constructor(private grpcClient: GRPCClient, private router: Router) {
  }

  add() {
    this.loading = true;
    if (this.url.startsWith("https://fabelio.com/")) {
      this.grpcClient.addProduct(this.url).then(res => {
        this.loading = false;
        this.router.navigateByUrl('/all');
      })
    } else {
      alert("The provided url not a fabelio product")
      this.loading = false;
      this.url = ""
    }

  }

}

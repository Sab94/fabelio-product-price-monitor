import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {AddComponent} from './components/add/add.component';
import {AllComponent} from './components/all/all.component';
import {DetailComponent} from './components/detail/detail.component';

const routes: Routes = [
  {
    path: 'add',
    component: AddComponent
  },
  {
    path: 'all',
    component: AllComponent
  },
  {
    path: 'detail/:id',
    component: DetailComponent
  },
  {
    path: '**',
    component: AllComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

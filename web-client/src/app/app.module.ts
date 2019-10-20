import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {AddComponent} from './components/add/add.component';
import { AllComponent } from './components/all/all.component';
import { DetailComponent } from './components/detail/detail.component';
import {FormsModule} from '@angular/forms';
import {NgxChartsModule} from '@swimlane/ngx-charts';
@NgModule({
  declarations: [
    AppComponent,
    AddComponent,
    AllComponent,
    DetailComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    NgxChartsModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

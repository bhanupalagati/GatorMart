import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CardComponent } from './card/card.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
// Material Imports
import {MatCardModule} from '@angular/material/card';
import {MatButtonModule} from '@angular/material/button';
import { ProductListComponent } from './product-list/product-list.component';
import { ProductDetailsComponent } from './product-details/product-details.component';
import {MatIconModule} from '@angular/material/icon';

// Other Imports
import {NgbPaginationModule, NgbAlertModule} from '@ng-bootstrap/ng-bootstrap';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FormsModule } from '@angular/forms';
import {HttpClientModule } from '@angular/common/http';

const materialImports = [
  MatCardModule,
  MatButtonModule,
  MatIconModule
]
@NgModule({
  declarations: [
    AppComponent,
    CardComponent,
    ProductListComponent,
    ProductDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ...materialImports,
    NgbModule,
    NgbPaginationModule,
    NgbAlertModule,
    FormsModule
  ],
  exports: [
    ...materialImports
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

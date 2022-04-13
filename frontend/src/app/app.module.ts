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
import {MatFormFieldModule} from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import {MatDialogModule} from '@angular/material/dialog';
import {MatMenuModule} from '@angular/material/menu';
import {MatChipsModule} from '@angular/material/chips';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
// Other Imports
import {NgbPaginationModule, NgbAlertModule} from '@ng-bootstrap/ng-bootstrap';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { LoaderComponent } from './loader/loader.component';
import { AdFormComponent } from './ad-form/ad-form.component';
import { DeletePopUpComponent } from './delete-pop-up/delete-pop-up.component';
import { CreateAddComponent } from './create-add/create-add.component';
import { EditAddComponent } from './edit-add/edit-add.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { InterceptorService } from './interceptor.service';
import { GoogleMapsModule } from '@angular/google-maps';
import { FiltersComponent } from './filters/filters.component';
import { HeaderComponent } from './header/header.component';


const materialImports = [
  MatCardModule,
  MatButtonModule,
  MatIconModule,
  MatFormFieldModule,
  MatInputModule,
  MatDialogModule,
  MatMenuModule,
  MatChipsModule,
  MatSlideToggleModule
]
@NgModule({
  declarations: [
    AppComponent,
    CardComponent,
    ProductListComponent,
    ProductDetailsComponent,
    LoaderComponent,
    AdFormComponent,
    DeletePopUpComponent,
    CreateAddComponent,
    EditAddComponent,
    LoginComponent,
    SignupComponent,
    FiltersComponent,
    HeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    GoogleMapsModule,
    ...materialImports,
    NgbModule,
    NgbPaginationModule,
    NgbAlertModule,
    FormsModule,
    ReactiveFormsModule
  ],
  exports: [
    ...materialImports
  ],
  providers: [{provide: HTTP_INTERCEPTORS, useClass: InterceptorService, multi: true}],
  bootstrap: [AppComponent]
})
export class AppModule { }

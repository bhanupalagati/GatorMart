import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateAddComponent } from './create-add/create-add.component';
import { EditAddComponent } from './edit-add/edit-add.component';
import { LoginComponent } from './login/login.component';
import { MyAdsComponent } from './my-ads/my-ads.component';
import { ProductDetailsComponent } from './product-details/product-details.component';
import { ProductListComponent } from './product-list/product-list.component';
import { SignupComponent } from './signup/signup.component';

const routes: Routes = [
  {path: '', component: LoginComponent},
  {path: 'signup', component: SignupComponent},
  {path: 'products', component: ProductListComponent},
  {path: 'create', component: CreateAddComponent},
  {path: 'edit', component: EditAddComponent},
  {path: 'myads', component: MyAdsComponent},
  {path: 'product/:id', component: ProductDetailsComponent},
  {path: 'deleteproduct', component: ProductDetailsComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

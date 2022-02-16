import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdFormComponent } from './ad-form/ad-form.component';
import { ProductDetailsComponent } from './product-details/product-details.component';
import { ProductListComponent } from './product-list/product-list.component';

const routes: Routes = [
  {path: '', component: ProductListComponent},
  {path: 'login', component: ProductListComponent},
  {path: 'signup', component: ProductListComponent},
  {path: 'create', component: AdFormComponent},
  {path: 'product/:id', component: ProductDetailsComponent},
  {path: 'createproduct', component: ProductDetailsComponent},
  {path: 'deleteproduct', component: ProductDetailsComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

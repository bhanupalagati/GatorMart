import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdFormComponent } from './ad-form/ad-form.component';
import { CreateAddComponent } from './create-add/create-add.component';
import { EditAddComponent } from './edit-add/edit-add.component';
import { ProductDetailsComponent } from './product-details/product-details.component';
import { ProductListComponent } from './product-list/product-list.component';

const routes: Routes = [
  {path: '', component: ProductListComponent},
  {path: 'login', component: ProductListComponent},
  {path: 'signup', component: ProductListComponent},
  {path: 'create', component: CreateAddComponent},
  {path: 'edit', component: EditAddComponent},
  {path: 'product/:id', component: ProductDetailsComponent},
  {path: 'deleteproduct', component: ProductDetailsComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

import { Injectable } from '@angular/core';
import { Product } from '../interfaces/product.interface';

@Injectable({
  providedIn: 'root'
})
export class EditFormService {
  form: Product;
  formGet;
  constructor() { }

  getFormData() {
    return this.formGet;
  }

  getProductID() {
    return this.form.ID;
  }

  setFormData(data) {
    this.form = { ...data };
    this.formGet = { title: this.form.title,
      secondary_title: this.form.secondary_title,
      imageUrl: this.form.imageUrl, price: this.form.price,
      images: this.form.images,
      simple_desc: this.form.simple_desc,
      description: this.form.description,
      city: this.form.city, state: this.form.state,
      location_lat: this.form.location_lat,
      location_long: this.form.location_long,
      condition: this.form.condition, age: this.form.age,
      target: this.form.target, category: this.form.category }
  }
}

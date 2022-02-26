import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Product, ProductResponse } from '../interfaces/product.interface';

@Injectable({
  providedIn: 'root'
})
export class ProductsService {
  baseUrl = "http://127.0.0.1:8000/";
  constructor(private http: HttpClient) { }

  getProductList() {
    return this.http.get<Product[]>(this.baseUrl + "products");
  }

  getProductDetails(id) {
    return this.http.get<Product>(this.baseUrl + "product/"+id);
  }

  createNewProduct(data) {
    return this.http.post(this.baseUrl + "product", data);
  }

  deleteProduct(id) {
    return this.http.delete(this.baseUrl + "product/"+id)
  }

  uploadImages(imageData) {
    return this.http.post(this.baseUrl + "product/upload", imageData)
  }
}

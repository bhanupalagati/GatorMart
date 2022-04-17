import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Product, ProductResponse } from '../interfaces/product.interface';
import { BehaviorSubject, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProductsService {
  baseUrl = "http://127.0.0.1:8000/";
  userData: any;
  filtersApplied = false;
  applyFilters = new BehaviorSubject({});
  dropDowns = {};
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

  updateProduct(id, data) {
    return this.http.put(this.baseUrl + "product/"+ id, data);
  }

  deleteProduct(id) {
    return this.http.delete(this.baseUrl + "product/"+id)
  }

  uploadImages(imageData) {
    return this.http.post(this.baseUrl + "product/upload", imageData)
  }

  signUpUser(userData) {
    return this.http.post(this.baseUrl+'register', userData);
  }

  signInUser(userData) {
    return this.http.post(this.baseUrl+'login', userData);
  }

  getFilteredProducts(filterData) {
    return this.http.post<Product[]>(this.baseUrl+'filterproducts', filterData);
  }

  getDropdown(name) {
    return this.http.get<string[]>(this.baseUrl+name);
  }

  setCookies(token) {
    let d = new Date();
    d.setTime(d.getTime()+(24*60*60*1000));

    let expires = "; expires="+d.toUTCString();
    document.cookie = "token="+token+expires+"; path=/"
  }

  getCookie(name='token') {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length >= 2) return parts.pop().split(";").shift();
  }

  validateCookie(token) {
    return this.http.post(this.baseUrl+'validate', token)
  }

  getUserData() {
    return this.userData;
  }

  setUserData(userData) {
    this.userData = JSON.parse(JSON.stringify(userData));
  }
}

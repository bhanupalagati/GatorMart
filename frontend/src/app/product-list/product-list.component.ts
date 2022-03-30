import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Product } from '../interfaces/product.interface';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})

export class ProductListComponent implements OnInit, OnDestroy {
  productResp: Product[];
  loading = true;
  subscription: Subscription;
  constructor(private productsService: ProductsService) { }

  ngOnInit(): void {
    this.productsService.getProductList().subscribe(res => {
      this.productResp = res;
      this.loading = false;
    });
    this.subscription = this.productsService.applyFilters.subscribe(filterData => {
      const filters = this.processFilterInfo(filterData);
      // Change this to fetch by filter route
      this.productsService.getProductList().subscribe(res => {
        this.productResp = res;
        this.loading = false;
      });
    });
  }

  processFilterInfo(filterData) {
    const prices = filterData['price'].split(" ");
    const age = filterData['age'].split(" ");
    if (prices.length > 2) {
      filterData['price'] = {
        operator: prices[0],
        value1: +prices[1],
        value2: +prices[2]
      }
    }
    if (age.length > 2) {
      filterData['age'] = {
        operator: age[0],
        value1: +age[1],
        value2: +age[2]
      }
    }
  
    return filterData;
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
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
  showFilter = false;
  productResp: Product[];
  loading = true;
  subscription: Subscription;
  constructor(private productsService: ProductsService) { }

  ngOnInit(): void {
    this.productsService.filterMenuToggled.subscribe(res => {
      this.showFilter = res;
    });
    
    this.apiCall(this.productsService.applyFilters.value);
    this.filterSubscription();
  }

  filterSubscription() {
    this.subscription = this.productsService.applyFilters.subscribe(filterData => {
      const filters = this.processFilterInfo(filterData);
      this.apiCall(filters)
    });
  }

  apiCall(filters) {
    this.productsService.getFilteredProducts(filters).subscribe(res => {
      this.productResp = res;
      this.loading = false;
    });
  }

  processFilterInfo(filterData) {
    const prices = filterData['price'] ? filterData['price'].split(" ") : [];
    const age = filterData['age'] ? filterData['age'].split(" ") : [];
    filterData['price'] = {};
    filterData['age'] = {};
    if (prices.length > 2) {
      filterData['price'] = {
        operator: prices[0],
        value1: +prices[1],
        value2: +prices[2],
        price: filterData.price
      }
    }
    if (age.length > 2) {
      filterData['age'] = {
        operator: age[0],
        value1: +age[1],
        value2: +age[2],
        age: filterData.age
      }
    }
  
    return filterData;
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}
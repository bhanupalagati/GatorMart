import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-filters',
  templateUrl: './filters.component.html',
  styleUrls: ['./filters.component.scss']
})
export class FiltersComponent implements OnInit {
  categories: string[];
  targets: string[];
  conditions: string[];
  sortBy = ['date', 'price']
  filterForm: FormGroup = new FormGroup({
    title: new FormControl(""),
    condition: new FormControl(""),
    target: new FormControl(""),
    age: new FormControl(""),
    sortBy: new FormControl(""),
    price: new FormControl(""),
    radius: new FormControl(""),
    category: new FormControl(""),
  });
  constructor(private productsService: ProductsService, private router: Router) { }

  ngOnInit(): void {
    this.getDropDownValues();
    if(this.productsService.filtersApplied) {
      const filterData = this.productsService.applyFilters.value
      this.filterForm.setValue({...filterData, price: filterData['price']['price'] || filterData['price'], age: filterData['age']['age'] || filterData['age']});
    }
  }

  getDropDownValues() {
    this.productsService.getDropdown('categories').subscribe(res => this.categories = res);
    this.productsService.getDropdown('target').subscribe(res => this.targets = res);
    this.productsService.getDropdown('condition').subscribe(res => this.conditions = res);
  }


  applyFilters() {
    this.productsService.filtersApplied = true;
    this.productsService.applyFilters.next(this.filterForm.value);
  }

  clearFilters() {
    this.productsService.filtersApplied = false;
    const clearedFilters = {title: "", condition: "", target: "", age: "", sortBy: "", price: "", radius: "", category: ""}
    this.productsService.applyFilters.next({...clearedFilters});
    this.filterForm.setValue({...clearedFilters});
  }

}
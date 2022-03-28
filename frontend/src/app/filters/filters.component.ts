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
    if(this.productsService.filtersApplied) {
      this.filterForm.setValue(this.productsService.applyFilters.value);
    }
  }

  applyFilters() {
    this.productsService.filtersApplied = true;
    this.productsService.applyFilters.next(this.filterForm.value);
  }

  clearFilters() {
    this.productsService.filtersApplied = false;
    this.productsService.applyFilters.next({});
  }

}
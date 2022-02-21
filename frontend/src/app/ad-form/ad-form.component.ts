import { LogicalFileSystem } from '@angular/compiler-cli/src/ngtsc/file_system';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import {MAT_FORM_FIELD, MatFormField, MatFormFieldControl} from '@angular/material/form-field';
import { ProductsService } from '../services/products.service';
@Component({
  selector: 'app-ad-form',
  templateUrl: './ad-form.component.html',
  styleUrls: ['./ad-form.component.scss']
})
export class AdFormComponent implements OnInit {
  createForm: FormGroup = new FormGroup({
    title: new FormControl("", [Validators.required]),
    secondary_title: new FormControl("", [Validators.required]),
    imageUrl: new FormControl("", [Validators.required]),
    price: new FormControl(0, [Validators.required]),
    simple_desc: new FormControl("", [Validators.required]),
    description: new FormControl("", [Validators.required]),
    city: new FormControl("", [Validators.required]),
    state: new FormControl("", [Validators.required]),
    location_lat: new FormControl("", [Validators.required]),
    location_long: new FormControl("", [Validators.required]),
    condition: new FormControl("NA", [Validators.required]),
    age: new FormControl(0, [Validators.required]),
    images: new FormControl("", [Validators.required]),
    target: new FormControl("", [Validators.required]),
    category: new FormControl("", [Validators.required])
  });
  constructor(private productsService: ProductsService) { }

  ngOnInit(): void {
  }

  createAd() {
    this.productsService.createNewProduct({...this.createForm.value, price: +this.createForm.value.price, age: +this.createForm.value.age}).subscribe(res => {});
  }

}

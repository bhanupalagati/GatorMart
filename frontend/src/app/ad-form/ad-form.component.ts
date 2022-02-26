import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ProductsService } from '../services/products.service';
@Component({
  selector: 'app-ad-form',
  templateUrl: './ad-form.component.html',
  styleUrls: ['./ad-form.component.scss']
})
export class AdFormComponent implements OnInit {
  formData = new FormData();
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
  constructor(private productsService: ProductsService, private router: Router) { }

  ngOnInit(): void {
  }

  createAd() {
    this.productsService.createNewProduct({...this.createForm.value, price: +this.createForm.value.price, age: +this.createForm.value.age}).subscribe(res => {
      this.router.navigate(['/']);
    });
  }

  onFileSelected(event, update) {
    this.formData = new FormData();
    if (event) {
      for (let i = 0; i < event.target.files.length; i++) {
        this.formData.append("photo", event.target.files[i]);
      }
      this.uploadAttachments(update);
    }
  }

  uploadAttachments(update) {
    this.productsService.uploadImages(this.formData).subscribe((res:any) => {
      this.createForm.controls[update].setValue(res.data.imageUrl.join());
      console.log(this.createForm.value);
    })
  }

}

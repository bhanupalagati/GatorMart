import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ProductsService } from '../services/products.service';
import { EditFormService } from './edit-form.service';
@Component({
  selector: 'app-ad-form',
  templateUrl: './ad-form.component.html',
  styleUrls: ['./ad-form.component.scss']
})
export class AdFormComponent implements OnInit {
  @Input() request;
  formData = new FormData();
  lati: string;
  longi: string;
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
  constructor(private productsService: ProductsService, private router: Router, private editFormService: EditFormService) { }

  ngOnInit(): void {
    if (this.request === 'edit') {
      this.populateEditForm();
    }
  }

  populateEditForm() {    
    this.createForm.setValue({...this.editFormService.getFormData()});
  }
  submitAdd() {
    if (this.request === 'create') {
      this.createAd();
    } else {
      this.updateAd();
    }
  }
  createAd() {
    this.productsService.createNewProduct({...this.createForm.value, price: +this.createForm.value.price, age: +this.createForm.value.age}).subscribe(res => {
      this.router.navigate(['/']);
    });
  }

  updateAd() {
    this.productsService.updateProduct(this.editFormService.getProductID(), {...this.createForm.value, price: +this.createForm.value.price, age: +this.createForm.value.age}).subscribe(res => {
      this.router.navigate(['/']);
    });
  }

  locateMe() {
    navigator.geolocation.getCurrentPosition((position) => {
      this.lati = position.coords.latitude.toString();
      this.longi = position.coords.longitude.toString();
      console.log('Geolocation test in progress:', this.lati, ' and ', this.longi)
      this.createForm.controls['location_lat'].setValue(this.lati);
      this.createForm.controls['location_long'].setValue(this.longi);
      this.getlocation(this.lati, this.longi).then((value) => {

        var city = value['address_components'][2]['long_name'];
        var state = value['address_components'][4]['long_name'];
        this.createForm.controls['city'].setValue(city);
        this.createForm.controls['state'].setValue(state);

      }).catch(console.error);
      
    }) 
  }

  getlocation(latitude, longitude) {
    return new Promise(function (resolve, reject) {
      var request = new XMLHttpRequest();
      var method = 'GET';
      var url = 'https://maps.googleapis.com/maps/api/geocode/json?latlng=' + latitude + ',' + longitude + atob('JmtleT1BSXphU3lDcWZaN2otTEZMdnZwRUxlX045bGtzcGVUVW9kX3NQOWs=');
      var async = true;

      request.open(method, url, async);
      request.onreadystatechange = function () {localStorage
          if (request.readyState == 4) {
              if (request.status == 200) {
                  var data = JSON.parse(request.responseText);
                  var address = data.results[0];
                  resolve(address);
              }
              else {
                  reject(request.status);
              }
          }
      };
      request.send();
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
    })
  }

}
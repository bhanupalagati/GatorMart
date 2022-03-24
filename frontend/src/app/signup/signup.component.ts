import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  signUpForm: FormGroup = new FormGroup({
    email: new FormControl("", [Validators.email]),
    password: new FormControl("", [Validators.required]),
    password2: new FormControl("", [Validators.required]),
    profession: new FormControl("", [Validators.required]),
    firstname: new FormControl("", [Validators.required]),
    lastname: new FormControl("", [Validators.required]),
    DOB: new FormControl("", [Validators.required]),
  });
  constructor() { }

  ngOnInit(): void {
  }

  checkCookies() {
    // If there was a previous login use that
  }
  signup() {
    // Make an API call if login is success
    // Save cookies
    // redirect to products
    
  }

}

import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup = new FormGroup({
    email: new FormControl("", [Validators.email]),
    password: new FormControl("", [Validators.required])
  });
  constructor() { }

  ngOnInit(): void {
  }

  checkCookies() {
    // If there was a previous login use that
  }
  login() {
    // Make an API call if login is success
    // Save cookies
    // redirect to products
  }

}

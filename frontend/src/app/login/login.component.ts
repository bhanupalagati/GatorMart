import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ProductsService } from '../services/products.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  token: string;
  loginForm: FormGroup = new FormGroup({
    email: new FormControl("", [Validators.email]),
    password: new FormControl("", [Validators.required])
  });
  constructor(private productsService: ProductsService, private router: Router) { }

  ngOnInit(): void {
    // this.checkCookies();
  }

  checkCookies() {
    this.token = this.productsService.getCookie();
    this.productsService.validateCookie({token: this.token}).subscribe(res => {

    });
  }
  login() {
    this.productsService.signInUser(this.loginForm.value).subscribe((res: any) => {
      this.productsService.setUserData(res);
      this.productsService.setCookies(res.token);
      this.router.navigate(['/products']);
  });
  }

}

import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { of } from 'rxjs';
import { CardComponent } from '../card/card.component';
import { productResp } from '../mocks/products.mock';
import { ProductsService } from '../services/products.service';

import { ProductListComponent } from './product-list.component';

describe('ProductListComponent', () => {
  let component: ProductListComponent;
  let fixture: ComponentFixture<ProductListComponent>;
  const productsServiceMock = {
    getProductList: () => {
      return of(productResp.products)
    }
  }
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProductListComponent ],
      providers: [{provide: ProductsService, useValue: productsServiceMock}],
      schemas: [CUSTOM_ELEMENTS_SCHEMA]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('Check access to the component', () => {
    expect(component).toBeTruthy();
  });

  //check if product count is 5git ch
  it('Check lenght of response', () => {
    console.log(component.productResp.length);
    expect(component.productResp.length).toEqual(8);
  });

  //check if product title is correct
  it('Check 1st products title', () => {
    console.log(component.productResp[1].title);
    expect(component.productResp[1].title).toEqual('Iphone 12 Pro Max');
  });

  //check if same number of html app cards
  it('should create', () => {
    console.log(fixture.nativeElement.querySelectorAll('app-card').length);
    expect(fixture.nativeElement.querySelectorAll('app-card').length).toEqual(8);
  });
});


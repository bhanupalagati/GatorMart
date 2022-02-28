import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
//import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';
import { product } from '../mocks/products.mock';
import { CardComponent } from './card.component';

describe('CardComponent', () => {
  let component: CardComponent;
  let fixture: ComponentFixture<CardComponent>;
  //creating a router mock and but a observable
  //is not required to be created here
  const routerMock  = {
    navigateByUrl: () => {},
  }
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CardComponent ],
      //injecting the router dependency here for the router
      //which is initialised in the constructor 
      providers: [{provide: Router, useValue: routerMock}],
      schemas: [CUSTOM_ELEMENTS_SCHEMA]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CardComponent);
    component = fixture.componentInstance;
    component.product = product;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
   // spyOn (TestBed.get(ProductsService), am).and.returnValue(false)
  });
});


//observabales
//behaviorals
 //RXJS package

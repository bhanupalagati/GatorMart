/// <reference types ="cypress" />

// Objective: Filter validation and result verification

describe('Login into the application', function () {

    it('Sign up functionality', function() {
        cy.visit('http://localhost:4200/')
        
    })
    it('Login function', function() {
        //Email
        cy.get('input')
        .eq(0).type('user@gatormart.com') 
        //Password
        cy.get('input')
        .eq(1).type('Password1$') 
    })
    it('Submit Button', function() {
        cy.get('button').click()       
    })
    it('Filter Test Check', function() {
        cy.visit('http://localhost:4200/products')
    })
    it('Product Filters', function() {
        //Title
        cy.get('input')
        .eq(0).type('iPhone') 
        //Condition
        cy.get('input')
        .eq(1).type('Used') 
        //Target
        cy.get('input')
        .eq(2).type('Student') 
        //Age
        cy.get('input')
        .eq(3).type('24') 
        //SortBy
        cy.get('input')
        .eq(4).type('name') 
        //Price
        cy.get('input')
        .eq(5).type('2300') 
        //Radius
        cy.get('input')
        .eq(6).type('2') 
        //Category
        cy.get('input')
        .eq(7).type('Electronics') 
    })
    it('Apply Button', function() {
        cy.get('button').eq(1).click()       
    })
    it('Product Filters', function() {
        //Title
        cy.get('input')
        .eq(0).type('Macbook') 
        //Condition
        cy.get('input')
        .eq(1).type('Used') 
        //Target
        cy.get('input')
        .eq(2).type('Wholesale') 
        //Age
        cy.get('input')
        .eq(3).type('24') 
        //SortBy
        cy.get('input')
        .eq(4).type('name') 
        //Price
        cy.get('input')
        .eq(5).type('2300') 
        //Radius
        cy.get('input')
        .eq(6).type('2') 
        //Category
        cy.get('input')
        .eq(7).type('Electronics') 
    })
    it('Apply Button', function() {
        cy.get('button').eq(0).click()       
    })
})
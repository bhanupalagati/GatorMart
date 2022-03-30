/// <reference types="cypress" />

//Test objective: 
//Accessing the website, check if the backend Api's are properly accessible.
//1. Access the main homepage
//2. Create the listing page
//3. Access a product in the website

describe('Access Test', function () {
    it('Hitting the backend', function() {
        cy.visit('http://localhost:4200/')
        cy.wait(2000)
        
    })
    it('Hitting the backend create', function() {
        cy.visit('http://localhost:4200/create')
        cy.wait(2000)
    })
    it('Hitting the backend create', function() {
        cy.visit('http://localhost:4200/product/23')
        cy.wait(2000)
    })
})


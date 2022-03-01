/// <reference types="cypress" />

//Accessing the website
describe('Access Test', function () {
    it('Hitting the backend', function() {
        cy.visit('http://localhost:4200/')
        
    })
})


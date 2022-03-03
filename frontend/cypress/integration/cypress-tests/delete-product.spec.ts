/// <reference types="cypress" />

describe('Deleting a product', function () {
    it('Access the page', function() {
        cy.visit('http://localhost:4200/')
        
    })

    it('get the first product and delete it', function() {
        cy.get('mat-card')
        .eq(0)      
        .click()
    })

    it('scroll to the remove button', function() {
        cy.scrollTo('bottom')
    })

    it('Click the Remove da button', function() {
        cy.get('button')
        .eq(1)
        .click()
    })
    
    //Use the fourth button, second button
    //will only reload the delete ad pop up button
    it('Delete the ad', function() {
        cy.get('button')
        .eq(3)
        .click()
    })
})
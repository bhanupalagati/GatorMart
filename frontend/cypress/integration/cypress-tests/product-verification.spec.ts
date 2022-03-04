/// <reference types="cypress" />

describe('Access the website', function () {
    beforeEach(() => {
        cy.visit('http://localhost:4200/')
    })
    it('Check the first product and verify the details ', function() {
        cy.get('mat-card').click()
        cy.wait(1500)
        cy.scrollTo('bottom')
        cy.wait(1500)
        cy.scrollTo('top') 
        
    })
})
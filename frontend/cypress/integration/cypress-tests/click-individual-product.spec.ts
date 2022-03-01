/// <reference types="cypress" />

describe('Access Test', function () {
    //Accessing the list of products:
    beforeEach(() => {
        cy.visit('http://localhost:4200/')
    })
    it('Click 1st product and move into detailed view', function() {
        cy.get('mat-card').click()
        cy.wait(1500)
        cy.scrollTo('bottom')
        cy.wait(1500)
        cy.scrollTo('top') 
        
    })
})
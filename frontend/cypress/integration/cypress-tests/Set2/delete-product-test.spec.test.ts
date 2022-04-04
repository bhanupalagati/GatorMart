/// <reference types="cypress" />

//Test objective: Check the edit functionality in the application

describe('Login into the app', function () {

        it('Login functionality', function() {
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

        it('Survey the product', function() {
            cy.get('mat-card').eq(1).click()
            cy.wait(1500)
            cy.scrollTo('bottom')
            cy.wait(1500)
            cy.scrollTo('top') 
        
    })
    it('Delete Product Button', function() {
        cy.get('button').eq(1).click()       
    })
    
    it('Verify Delete function', function() {
        cy.get('button').eq(14).click() 
    })


})
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
    it('Edit Button', function() {
        cy.get('button').eq(0).click()       
    })
    
    it('Enter the edited details', function() {
        //Price
         cy.get('input')
         .eq(5)
         .clear()
         .type('25500')
        //Simple Description
        cy.get('input')
         .eq(6)
         .clear()
         .type('Edited ad description')        
    })

    
    it('Submit the form', function() {
        cy.get('button').eq(1).click()       
    })

    


})
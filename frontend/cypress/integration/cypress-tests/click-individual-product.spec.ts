/// <reference types="cypress" />

describe('Access Test', function () {
    //Accessing the list of products:
    beforeEach(() => {
        cy.visit('http://localhost:4200/')
    })

    it('Click 1st product and move into detailed view', function() {
        cy.get('mat-card').eq(1).click()
        cy.wait(1500)
        cy.scrollTo('bottom')
        cy.wait(1500)
        cy.scrollTo('top') 
        
  //  })
    //Also check the carousel
  //  it('Control the carousel movement', function() {
        cy.get('app-product-details')
        .find('ngb-carousel')
        .find('a')
        .eq(1)
        .click()
        cy.wait(1000)
        cy.get('app-product-details')
        .find('ngb-carousel').find('a').eq(1)
        .click()
        cy.wait(1000)
        cy.get('app-product-details')
        .find('ngb-carousel').find('a')
        .eq(1)
        .click()
        cy.wait(1000)
        cy.get('app-product-details')
        .find('ngb-carousel')
        .find('a').eq(0)
        .click()
        cy.wait(1000)
        cy.get('app-product-details')
        .find('ngb-carousel').find('a').eq(0)
        .click()
        cy.wait(1000)
        
    })
})
/**
 * Copyright 2014 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2013-12-03 23:20
 * description :
 * history :
 */

package dps

import (
	"github.com/atnet/gof"
	"go2o/src/core/query"
	"go2o/src/core/repository"
)

var (
	Context         gof.App
	PromService     *promotionService
	ShoppingService *shoppingService
	MemberService   *memberService
	PartnerService  *partnerService
	SaleService     *saleService
	DeliverService  *deliveryService
)

func Init(ctx gof.App) {
	Context := ctx
	db := Context.Db()

	/** Repository **/
	userRep := repository.NewUserRep(db)
	memberRep := repository.NewMemberRep(db)
	partnerRep := repository.NewPartnerRep(db, userRep, memberRep)
	promRep := repository.NewPromotionRep(db, memberRep)
	tagSaleRep := repository.NewTagSaleRep(db)
	saleRep := repository.NewSaleRep(db, tagSaleRep)
	deliveryRep := repository.NewDeliverRep(db)
	spRep := repository.NewShoppingRep(db, partnerRep, saleRep, promRep, memberRep, deliveryRep)

	/** Query **/
	mq := query.NewMemberQuery(db)
	pq := query.NewPartnerQuery(ctx)

	/** Service **/
	PromService = NewPromotionService(promRep)
	ShoppingService = NewShoppingService(spRep)
	MemberService = NewMemberService(memberRep, mq)
	PartnerService = NewPartnerService(partnerRep, pq)
	SaleService = NewSaleService(saleRep)
	DeliverService = NewDeliveryService(deliveryRep)
}

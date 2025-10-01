package dummydata

import (
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

var Categories = []models.Category{
	{Name: "Smoothies"},
	{Name: "Pasta"},
	{Name: "Oats"},
}

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := SeedCategories(tx); err != nil {
			return err
		}
		if err := SeedProducts(tx); err != nil {
			return err
		}
		return nil
	})
}

func SeedCategories(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, c := range Categories {
			var existing models.Category
			if err := tx.First(&existing, "id = ?", c.ID).Error; err == gorm.ErrRecordNotFound {
				if err := tx.Create(&c).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func SeedProducts(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, p := range Products {
			var existing models.Product
			if err := tx.First(&existing, "id = ?", p.ID).Error; err == gorm.ErrRecordNotFound {
				if err := tx.Create(&p).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

var Products = []models.Product{
	{
		ID:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Mango + Greens",
		Description: "This creamy smoothie gets its vivid hue from nutrient-dense spinach and its lush flavor from juicy mango, coconut cream, and a spritz of lime. Oh, and we added kohlrabi because we like you, and it's a bit of a nutritional powerhouse.",
		SKU:         "47175570784554:0ce54760f167fb2d13aa9f3761758d22",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A10-MANGO.jpg?v=1716498612",
		CategoryID:  1,
	},
	{
		ID:          2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Tart Cherry + Raspberry",
		Description: "Simply the berriest berry smoothie we could dream up. Raspberries, blueberries, and cherries (ok, technically not a berry) are brilliantly hued and full of antioxidants and anti-inflammatory properties. The natural sweetness of mission figs complements all that juicy tartness.",
		SKU:         "46783280054570:355fa0a7f2ff85196209c355d27c0b63",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A25-RAZZ.jpg?v=1716504576",
		CategoryID:  1,
	},
	{
		ID:          3,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Ginger + Greens",
		Description: "The fiery zing of ginger. The velvety texture of bananas, avocado, and flax seeds. A hint of sweetness from dates. A squeeze of lemon. And, of course, nutrient-dense, bright green spinach. It's green juice (but with fiber!) in highly-sippable smoothie form.",
		SKU:         "47175581303082:a3d82fd61c16f9b00746cdf59aace8c6",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A17-GING.jpg?v=1716498612",
		CategoryID:  1,
	},
	{
		ID:          4,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Cherry + Almond",
		Description: "Imagine a nut butter & jelly sandwich: sweet but not too sweet with cherries, juicy strawberries, and a sprinkle of vanilla bean. Rich and satisfying with almond butter and toasty sacha inchi seeds. No crusts here; we're purists.",
		SKU:         "46783324029226:15df0d4cc386a5933eb6cacd5020b329",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A22-PBNJ.jpg?v=1716498612",
		CategoryID:  1,
	},
	{
		ID:          5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Mango + Papaya",
		Description: "A naturally sweet blend of hydrating pineapple, mango, and papaya. Some macadamias for that good fat. This smoothie captures the color—and, more importantly, the vibe—of a summer sunset.",
		SKU:         "46783308824874:e99155c047680879286a17d6238de47e",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A03-PAPA.jpg?v=1716498612",
		CategoryID:  1,
	},
	{
		ID:          6,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "White Bean + Spinach Pesto",
		Description: "Yes, this pasta is gluten-free. Yes, this pesto is nut- and dairy-free. And yes, it tastes like the bright, fresh, al dente summer dish we dream about all year long. Built on sunflower seed butter, basil, garlic, and lemon, with a dash of nutritional yeast for a parmesan-like hit of umami, our pesto is rich and creamy. It clings deliciously to sedani, a tubular pasta made of ancient grains and just right for catching the herby sauce. Artichokes and beans add flavor and fiber.",
		SKU:         "47121324802346:d44ee605a9e206dbe5b6d35a1743b03f",
		BasePrice:   107.00,
		Stock:       15,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/R05-GRBEAN.jpg?v=1716498567",
		CategoryID:  2,
	},
	{
		ID:          7,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Tomato Basil + Portobello Bolognese",
		Description: "When all you want is a bottomless bowl of bolognese, these saucy noodles do the trick. Rich tomatoes and onions coat gluten-free sedani—noodles built on ancient grains and perfect for catching sauce. Basil brings brightness, oregano adds herby flair, and portobellos and spinach deliver a hearty bite. Hemp seeds and black lentils add plenty of fiber.",
		SKU:         "47121324835114:a8e4f04f85470e56bd3ec82d9bf7a17d",
		BasePrice:   107.00,
		Stock:       15,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/R06-BOLO.jpg?v=1716498567",
		CategoryID:  2,
	},
	{
		ID:          8,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Banana + Almond",
		Description: "This one's for the banana bread fans. Think perfectly ripe bananas and vanilla beans topped with velvety, satisfying nut butter (a mix of hazelnuts, cashews, almonds, and sunflower seeds). Cacao and maca root add rich flavor and maybe a little boost of energy, too.",
		SKU:         "46783296045354:487875c22e6db910255951f1557e1f3c",
		BasePrice:   84.00,
		Stock:       12,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A28-SkippyBanana_Almond.png?v=1746553810",
		CategoryID:  1,
	},
	{
		ID:          9,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Cinnamon + Banana",
		Description: "Name a more perfectly balanced breakfast. We'll wait. Gluten-free rolled oats and sweet potato serve the sustained energy. Maple and cinnamon add a dash of sweet 'n spice, giving it that banana bread taste. The real secret ingredients though? Reishi mushrooms for a little added richness and a whole lot of bliss-boosting. Sounds pretty perfect to us.",
		SKU:         "47121677877546:1535ed89a97a1e5f777f1120bbf30956",
		BasePrice:   67.00,
		Stock:       8,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/D01-CINNA_9dc0c535-c0f0-4a64-9c61-b7d453f3cd60.jpg?v=1716504918",
		CategoryID:  3,
	},
	{
		ID:          10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Strawberry + Peach",
		Description: "If you asked a peach what it wanted to be when it grew up, it would tell you: THIS SMOOTHIE. Sweet strawberries, bright raspberries, and a hint of tartness from goji berries round out that irresistibly juicy peach flavor. Bananas, oats, and flax seeds make the whole thing creamy and satisfying.",
		SKU:         "46783262982442:8351e728d69a8ba6d991329fb0588c65",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/A04-STRAW.jpg?v=1716498612",
		CategoryID:  1,
	},
	{
		ID:          11,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Title:       "Apple + Cinnamon",
		Description: "As easy as apple pie—and even more wholesome. Spiced apples bring that cinnamony goodness, while coconut cream and toasted vanilla bean offer the deep, rich taste of a just-starting-to-melt scoop of French vanilla ice cream. With adaptogenic maca. This Smoothie serves up a little slice of perfection.",
		SKU:         "46783285264682:4d66c5484db846bd8090d24690a80272",
		BasePrice:   84.00,
		Stock:       10,
		IsActive:    true,
		ImageURL:    "https://cdn.shopify.com/s/files/1/0732/7567/0826/files/a19-cook-vanilla-bean-cacao-deconstructed-pdp-web-122824.jpg?v=1735398346",
		CategoryID:  1,
	},
	// Continue similarly for remaining products...
}

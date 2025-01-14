package unit

import (
	"github.com/Ramnarong06/SE_PreLabTest2_2/entity"
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func TestUnitUser(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run("Correct", func(t *testing.T){
		user := entity.User{
			Username:		"Ramnarong",		// ไม่ "" ที่ด้านหน้า
			Email:			"b6513214@gmail.com",
			Password:		"12345",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestUnitUsername(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run("Username is required", func(t *testing.T){
		user := entity.User{
			Username:		"",
			Email:			"b6513214@gmail.com",
			Password:		"12345",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Username is required"))  // ข้อความต้องตรงกับ Entity

	})
}

func TestUnitUseremail(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run("Email is Invalid", func(t *testing.T){
		user := entity.User{
			Username:	"Ramnarong",
			Email:		"b6513214",
			Password:	"12345",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Email is Invalid"))
	})
}
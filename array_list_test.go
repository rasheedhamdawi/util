package util_test

import (
	"testing"

	"github.com/rasheedhamdawi/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArrayListSuite struct {
	suite.Suite
	arrayList util.List
}

// before each test
func (suite *ArrayListSuite) SetupTest() {
	suite.arrayList = new(util.ArrayList)
}

func (suite *ArrayListSuite) TestSize() {
	assert.Equal(suite.T(), 0, suite.arrayList.Size())
}

func (suite *ArrayListSuite) TestIsEmpty() {
	assert.True(suite.T(), suite.arrayList.IsEmpty())
}

func (suite *ArrayListSuite) TestAdd() {
	assert.True(suite.T(), suite.arrayList.Add("Go"))
}

func (suite *ArrayListSuite) TestGet() {
	assert.Nil(suite.T(), suite.arrayList.Get(5))
	suite.arrayList.Add("Python")
	suite.arrayList.Add("Go")
	suite.arrayList.Add("Java")
	assert.Equal(suite.T(), "Go", suite.arrayList.Get(1))
}

func (suite *ArrayListSuite) TestContains() {
	suite.arrayList.Add("Go")
	assert.True(suite.T(), suite.arrayList.Contains("Go"))
	assert.False(suite.T(), suite.arrayList.Contains("Java"))
}

func (suite *ArrayListSuite) TestIndexOf() {
	assert.Equal(suite.T(), -1, suite.arrayList.IndexOf("JavaScript"))
	suite.arrayList.Add("Python")
	suite.arrayList.Add("Go")
	suite.arrayList.Add("Java")
	assert.Equal(suite.T(), 1, suite.arrayList.IndexOf("Go"))
}

func (suite *ArrayListSuite) TestClear() {
	suite.arrayList.Add("Python")
	suite.arrayList.Add("Go")
	suite.arrayList.Add("Java")
	assert.Equal(suite.T(), 3, suite.arrayList.Size())
	suite.arrayList.Clear()
	assert.Equal(suite.T(), 0, suite.arrayList.Size())
}

func (suite *ArrayListSuite) TestIterator() {
	suite.arrayList.Add("Python")
	suite.arrayList.Add("Go")
	suite.arrayList.Add("Java")

	iterator := suite.arrayList.Iterator()

	for iterator.HasNext() {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, iterator.Next())
	}

}

func (suite *ArrayListSuite) TestForEach() {
	suite.arrayList.Add("Python")
	suite.arrayList.Add("Go")
	suite.arrayList.Add("Java")

	suite.arrayList.ForEach(func(i int, v util.Element) {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, v)
	})
}

func (suite *ArrayListSuite) TestItems() {
	suite.arrayList.Add("Python")
	suite.arrayList.Add("Go")
	suite.arrayList.Add("Java")

	for val := range suite.arrayList.Items() {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, val)
	}
}

func TestArrayListSuite(t *testing.T) {
	suite.Run(t, new(ArrayListSuite))
}

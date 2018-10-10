package util_test

import (
	"testing"

	"github.com/rasheedhamdawi/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HashSetSuite struct {
	suite.Suite
	hashSet util.Set
}

// before each test
func (suite *HashSetSuite) SetupTest() {
	suite.hashSet = new(util.HashSet)
}

func (suite *HashSetSuite) TestSize() {
	assert.Equal(suite.T(), 0, suite.hashSet.Size())
}

func (suite *HashSetSuite) TestIsEmpty() {
	assert.True(suite.T(), suite.hashSet.IsEmpty())
}

func (suite *HashSetSuite) TestAdd() {
	assert.True(suite.T(), suite.hashSet.Add("Go"))
}

func (suite *HashSetSuite) TestContains() {
	suite.hashSet.Add("Go")
	assert.True(suite.T(), suite.hashSet.Contains("Go"))
	assert.False(suite.T(), suite.hashSet.Contains("Java"))
}

func (suite *HashSetSuite) TestClear() {
	suite.hashSet.Add("Python")
	suite.hashSet.Add("Go")
	suite.hashSet.Add("Java")
	assert.Equal(suite.T(), 3, suite.hashSet.Size())
	suite.hashSet.Clear()
	assert.Equal(suite.T(), 0, suite.hashSet.Size())
}

func (suite *HashSetSuite) TestRemove() {
	suite.hashSet.Add("Python")
	suite.hashSet.Add("Go")
	suite.hashSet.Add("Java")
	assert.Nil(suite.T(), suite.hashSet.Remove("Java"))
}

func (suite *HashSetSuite) TestToArray() {
	suite.hashSet.Add("Python")
	suite.hashSet.Add("Go")
	suite.hashSet.Add("Java")

	for _, item := range suite.hashSet.ToArray() {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, item)
	}
}

func (suite *HashSetSuite) TestIterator() {
	suite.hashSet.Add("Python")
	suite.hashSet.Add("Go")
	suite.hashSet.Add("Java")

	iterator := suite.hashSet.Iterator()

	for iterator.HasNext() {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, iterator.Next())
	}
}

func (suite *HashSetSuite) TestForEach() {
	suite.hashSet.Add("Python")
	suite.hashSet.Add("Go")
	suite.hashSet.Add("Java")

	suite.hashSet.ForEach(func(i int, v util.Element) {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, v)
	})
}

func (suite *HashSetSuite) TestItems() {

	suite.hashSet.Add("Python")
	suite.hashSet.Add("Go")
	suite.hashSet.Add("Java")

	for val := range suite.hashSet.Items() {
		assert.Contains(suite.T(), []string{"Go", "Java", "Python"}, val)
	}
}

func TestHashSetSuite(t *testing.T) {
	suite.Run(t, new(HashSetSuite))
}

package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{ID: 1, Title: "Test Article 1", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras at elit tortor. Nullam porttitor est vitae scelerisque sollicitudin. Mauris at fringilla ante, a maximus justo. Ut maximus erat nisi, consequat cursus diam fermentum tempus. Etiam sit amet dolor sit amet mauris dapibus interdum nec non massa."},
	{ID: 2, Title: "Test Article 2", Content: "Proin pulvinar nisl nec fermentum fermentum. Mauris sed lacus erat. Fusce nec est congue, vulputate nisi non, tincidunt quam. Etiam at pretium diam. Phasellus semper nibh gravida commodo lacinia. Curabitur lorem ligula, interdum at pulvinar in, fermentum in orci. Pellentesque eleifend ultricies arcu eu suscipit."},
}

func GetAllArticles() []Article {
	return articleList
}

func GetArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("article not found")
}

func CreateNewArticle(title, content string) (*Article, error) {
	a := Article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}

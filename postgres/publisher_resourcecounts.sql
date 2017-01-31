SELECT s.s_identifier AS identifier,s.type_title as title, count(*) AS count FROM (
  SELECT  resources.id as r_id,
          resources.type,
          resources.is_disabled,
          resources.state,
          resources.status,
          resources.publisher_id,
          resources.resource_type_id,
          standards.identifier AS s_identifier,
          alignments.standard_id AS s_id,
          categories.standard_group_id as standard_group_id,
          resource_types.title as type_title
  FROM resources
  INNER JOIN alignments
    ON alignments.resource_id = resources.id
  INNER JOIN standards
    ON alignments.standard_id = standards.id
  INNER JOIN categories
    ON standards.category_id = categories.id
      AND categories.standard_group_id IN (1,2,4,6)
  INNER JOIN resource_types ON resource_types.id = resources.resource_type_id
  WHERE resource_types.title IN ('Video','Game')
  AND alignments.status = 2
  AND resources.is_disabled IN (NULL, false)
  AND resources.state = 'public'
  AND resources.status is NULL
) s GROUP BY resources.publisher_id,s.type_title ORDER BY count desc
